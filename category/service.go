package category

import (
	"context"
	"github.com/uanderson/pockee/autoid"
	"github.com/uanderson/pockee/category/dao"
	"github.com/uanderson/pockee/database"
	"github.com/uanderson/pockee/echox"
	"github.com/uanderson/pockee/errorsx"
)

type Service struct {
	dao *dao.Queries
}

func NewService(database *database.Database) *Service {
	return &Service{dao: dao.New(database.Pool)}
}

func (s *Service) ExistsCategoryByID(ctx context.Context, ID string) (bool, error) {
	return s.dao.ExistsCategoryByID(ctx, dao.ExistsCategoryByIDParams{
		ID:     ID,
		UserID: echox.GetUserID(ctx),
	})
}

func (s *Service) GetCategories(ctx context.Context) ([]dao.Category, error) {
	categories, err := s.dao.GetCategories(ctx, echox.GetUserID(ctx))
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *Service) CreateCategory(ctx context.Context, input CreateCategoryInput) error {
	err := s.dao.CreateCategory(ctx, dao.CreateCategoryParams{
		ID:     autoid.New(),
		Name:   input.Name,
		UserID: echox.GetUserID(ctx),
	})

	return err
}

func (s *Service) UpdateCategory(ctx context.Context, input UpdateCategoryInput) error {
	exists, err := s.ExistsCategoryByID(ctx, input.ID)
	if err != nil {
		return err
	}

	if !exists {
		return errorsx.CategoryNotFound
	}

	return s.dao.UpdateCategory(ctx, dao.UpdateCategoryParams{
		ID:     input.ID,
		Name:   input.Name,
		UserID: echox.GetUserID(ctx),
	})
}

func (s *Service) DeleteCategory(ctx context.Context, input DeleteCategoryInput) error {
	exists, err := s.ExistsCategoryByID(ctx, input.ID)
	if err != nil {
		return err
	}

	if !exists {
		return errorsx.CategoryNotFound
	}

	return s.dao.DeleteCategory(ctx, dao.DeleteCategoryParams{
		ID:     input.ID,
		UserID: echox.GetUserID(ctx),
	})
}
