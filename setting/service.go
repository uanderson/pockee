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

type Service struct {
	dao *dao.Queries
}

func NewService(database *database.Database) *Service {
	return &Service{dao: dao.New(database.Pool)}
}

func (s *Service) GetSettingByKey(ctx context.Context, key string) (dao.Setting, error) {
	return s.dao.GetSettingByKey(ctx, key)
}

func (s *Service) GetUserSettingByKey(ctx context.Context, key string) (dao.UserSetting, error) {
	userID := ctx.Value("userID").(string)

	return s.dao.GetUserSettingByKey(context.Background(), dao.GetUserSettingByKeyParams{
		Key:    key,
		UserID: userID,
	})
}

func (s *Service) UpdateUserSetting(ctx context.Context, input *UpdateSettingInput) (dao.UserSetting, error) {
	if !isSettingKeyAllowed(input.Key) {
		return dao.UserSetting{}, validation.NewError(fmt.Sprintf("Key '%s' is not allowed", input.Key))
	}

	userID, err := util.GetUserID(ctx)
	if err != nil {
		return dao.UserSetting{}, err
	}

	err = s.dao.UpdateUserSetting(ctx, dao.UpdateUserSettingParams{
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

func isSettingKeyAllowed(key string) bool {
	allowedSettingsKeys := []string{"inter.mtls.cert", "inter.mtls.key"}

	for _, allowedSettingKey := range allowedSettingsKeys {
		if allowedSettingKey == key {
			return true
		}
	}

	return false
}
