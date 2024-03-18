package setting

import (
	"context"
	"github.com/uanderson/pockee/autoid"
	"github.com/uanderson/pockee/database"
	"github.com/uanderson/pockee/echox"
	"github.com/uanderson/pockee/errorsx"
	"github.com/uanderson/pockee/setting/dao"
)

type Service struct {
	dao *dao.Queries
}

func NewService(database *database.Database) *Service {
	return &Service{dao: dao.New(database.Pool)}
}

func (s *Service) ExistsSettingByKey(ctx context.Context, key string) (bool, error) {
	return s.dao.ExistsSettingByKey(ctx, dao.ExistsSettingByKeyParams{
		Key:    key,
		UserID: echox.GetUserID(ctx),
	})
}

func (s *Service) DeleteSetting(ctx context.Context, input DeleteSettingInput) error {
	exists, err := s.ExistsSettingByKey(ctx, input.Key)
	if err != nil {
		return err
	}

	if !exists {
		return errorsx.SettingNotFound
	}

	return s.dao.DeleteSettingByKey(ctx, dao.DeleteSettingByKeyParams{
		Key:    input.Key,
		UserID: echox.GetUserID(ctx),
	})
}

func (s *Service) UpdateUserSetting(ctx context.Context, input UpdateSettingInput) error {
	if !isSettingKeyAllowed(input.Key) {
		return errorsx.SettingNotFound
	}

	return s.dao.UpdateSetting(ctx, dao.UpdateSettingParams{
		ID:     autoid.New(),
		Key:    input.Key,
		Value:  input.Value,
		UserID: echox.GetUserID(ctx),
	})
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
