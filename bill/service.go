package bill

import (
	"github.com/uanderson/pockee/category/dao"
	"github.com/uanderson/pockee/database"
)

type Service struct {
	dao *dao.Queries
}

func NewService(database *database.Database) *Service {
	return &Service{dao: dao.New(database.Pool)}
}
