package setting

import (
	"context"
	"github.com/uanderson/pockee/database"
	"github.com/uanderson/pockee/setting/dao"
)

type Service struct {
	dao *dao.Queries
}

func (s *Service) GetSettingByKey(key string) (dao.Setting, error) {
	return s.dao.GetSettingByKey(context.Background(), key)
}

func NewService() Service {
	return Service{
		dao: dao.New(database.Pool),
	}
}
