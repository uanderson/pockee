package setting

import (
	"context"
	"fmt"
	"github.com/uanderson/pockee/autoid"
	"github.com/uanderson/pockee/database"
	"github.com/uanderson/pockee/setting/dao"
	"github.com/uanderson/pockee/util"
	"github.com/uanderson/pockee/validation"
)

var allowedSettingsKeys = []string{"inter.mtls.cert", "inter.mtls.key"}

// Service holds an instance of the dao
type Service struct {
	dao *dao.Queries
}

// NewService creates a new instance of Service
func NewService(database *database.Database) *Service {
	return &Service{dao: dao.New(database.Pool)}
}

// GetSettingByKey returns a setting owned by the system
func (s *Service) GetSettingByKey(key string) (dao.Setting, error) {
	return s.dao.GetSettingByKey(context.Background(), key)
}

// GetUserSettingByKey returns a setting owned by the user
func (s *Service) GetUserSettingByKey(ctx context.Context, key string) (dao.UserSetting, error) {
	userID := ctx.Value("userID").(string)

	return s.dao.GetUserSettingByKey(context.Background(), dao.GetUserSettingByKeyParams{
		Key:    key,
		UserID: userID,
	})
}

// UpdateUserSetting persists to the database the setting updates
func (s *Service) UpdateUserSetting(ctx context.Context, input *UpdateSettingInput) (dao.UserSetting, error) {
	if !isSettingKeyAllowed(input.Key, allowedSettingsKeys) {
		return dao.UserSetting{}, validation.NewError(fmt.Sprintf("Key '%s' is not allowed", input.Key))
	}

	userID := util.GetUserID(ctx)
	err := s.dao.UpdateUserSetting(ctx, dao.UpdateUserSettingParams{
		ID:     autoid.New(),
		Key:    input.Key,
		Value:  input.Value,
		UserID: userID,
	})

	if err != nil {
		return dao.UserSetting{}, err
	}

	return s.GetUserSettingByKey(ctx, input.Key)
}

func isSettingKeyAllowed(key string, allowedSettingsKeys []string) bool {
	for _, allowedSettingKey := range allowedSettingsKeys {
		if allowedSettingKey == key {
			return true
		}
	}

	return false
}
