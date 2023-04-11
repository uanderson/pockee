package setting

import (
	"context"
	"fmt"
	"github.com/uanderson/pockee/autoid"
	"github.com/uanderson/pockee/database"
	"github.com/uanderson/pockee/setting/dao"
	"github.com/uanderson/pockee/validation"
)

var allowedSettingsKeys = []string{"inter.mtls.cert", "inter.mtls.key"}

// Service holds an instance of the dao
type Service struct {
	dao *dao.Queries
}

// UpdateSettingInput holds information on which
// data should be updated for a specific setting
type UpdateSettingInput struct {
	Key   string `json:"key" validate:"required,max=255"`
	Value string `json:"value" validate:"required,max=65535"`
}

// NewService creates a new instance of Service
func NewService() *Service {
	return &Service{
		dao: dao.New(database.Pool),
	}
}

// GetSettingByKey returns a setting owned by the system
func (s *Service) GetSettingByKey(key string) (dao.Setting, error) {
	return s.dao.GetSettingByKey(context.Background(), key)
}

// GetUserSettingByKey returns a setting owned by the user
func (s *Service) GetUserSettingByKey(key string, userId string) (dao.UserSetting, error) {
	return s.dao.GetUserSettingByKey(context.Background(), dao.GetUserSettingByKeyParams{
		Key:    key,
		UserId: userId,
	})
}

// UpdateUserSetting persists to the database the setting updates
func (s *Service) UpdateUserSetting(input *UpdateSettingInput, userId string) (dao.UserSetting, error) {
	if !isSettingKeyAllowed(input.Key, allowedSettingsKeys) {
		return dao.UserSetting{}, validation.NewError(fmt.Sprintf("Key '%s' is not allowed", input.Key))
	}

	err := s.dao.UpdateUserSetting(context.Background(), dao.UpdateUserSettingParams{
		Id:     autoid.Id(),
		Key:    input.Key,
		Value:  input.Value,
		UserId: userId,
	})

	if err != nil {
		return dao.UserSetting{}, err
	}

	return s.GetUserSettingByKey(input.Key, userId)
}

// isSettingKeyAllowed verifies if the provided key is allowed
// in the allowedSettingsKeys
func isSettingKeyAllowed(key string, allowedSettingsKeys []string) bool {
	for _, allowedSettingKey := range allowedSettingsKeys {
		if allowedSettingKey == key {
			return true
		}
	}
	return false
}
