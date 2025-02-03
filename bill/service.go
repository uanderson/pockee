package bill

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/uanderson/pockee/autoid"
	"github.com/uanderson/pockee/bill/dao"
	"github.com/uanderson/pockee/category"
	"github.com/uanderson/pockee/contact"
	"github.com/uanderson/pockee/database"
	"github.com/uanderson/pockee/echox"
	"github.com/uanderson/pockee/errorsx"
	"github.com/uanderson/pockee/utils"
	"time"
)

type Service struct {
	dao             *dao.Queries
	categoryService *category.Service
	contactService  *contact.Service
	pool            *pgxpool.Pool
}

func NewService(database *database.Database, categoryService *category.Service, contactService *contact.Service) *Service {
	return &Service{
		dao:             dao.New(database.Pool),
		categoryService: categoryService,
		contactService:  contactService,
		pool:            database.Pool,
	}
}

func (s *Service) CreateBill(ctx context.Context, input CreateBill) error {
	existsCategory, err := s.categoryService.ExistsCategoryByID(ctx, input.CategoryID)
	if err != nil || !existsCategory {
		return errorsx.CategoryNotFound
	}

	existsContact, err := s.contactService.ExistsContactByID(ctx, input.ContactID)
	if err != nil || !existsContact {
		return errorsx.ContactNotFound
	}

	startAt, err := time.Parse("2006-01-02", input.DueAt)
	if err != nil || startAt.Before(time.Now()) {
		return errorsx.BillInvalidDueAt
	}

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	qtx := s.dao.WithTx(tx)

	billParams := dao.CreateBillParams{
		ID:          autoid.New(),
		Description: input.Description,
		Type:        input.Type,
		DueAt:       pgtype.Date{Time: startAt, Valid: true},
		Amount:      input.Amount,
		Status:      "PENDING",
		CategoryID:  input.CategoryID,
		ContactID:   input.ContactID,
		UserID:      echox.GetUserID(ctx),
	}

	err = qtx.CreateBill(ctx, billParams)
	if err != nil {
		return err
	}

	if input.Type == "BOLETO" {
		err = qtx.CreateBoleto(ctx, dao.CreateBoletoParams{
			ID:      autoid.New(),
			Barcode: utils.Deref(input.Barcode, ""),
			BillID:  billParams.ID,
		})
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func (s *Service) CreateRecurringBill(ctx context.Context, input CreateRecurringBill) error {
	existsCategory, err := s.categoryService.ExistsCategoryByID(ctx, input.CategoryID)
	if err != nil || !existsCategory {
		return errorsx.CategoryNotFound
	}

	existsContact, err := s.contactService.ExistsContactByID(ctx, input.ContactID)
	if err != nil || !existsContact {
		return errorsx.ContactNotFound
	}

	startAt, err := time.Parse("2006-01-02", input.StartAt)
	if err != nil || startAt.Before(time.Now()) {
		return errorsx.BillInvalidStartAt
	}

	params := dao.CreateRecurringBillParams{
		ID:          autoid.New(),
		Description: input.Description,
		Type:        input.Type,
		Amount:      input.Amount,
		Interval:    input.Interval,
		CategoryID:  input.CategoryID,
		ContactID:   input.ContactID,
		UserID:      echox.GetUserID(ctx),
		StartAt:     pgtype.Date{Time: startAt, Valid: true},
	}

	if input.EndAt != nil {
		endAt, err := time.Parse("2006-01-02", *input.EndAt)
		if err != nil || endAt.Before(time.Now()) || endAt.Before(startAt) {
			return errorsx.BillInvalidEndAt
		}

		params.EndAt = pgtype.Date{Time: endAt, Valid: true}
	}

	return s.dao.CreateRecurringBill(ctx, params)
}
