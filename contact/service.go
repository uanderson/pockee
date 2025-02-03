package contact

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/uanderson/pockee/autoid"
	"github.com/uanderson/pockee/contact/dao"
	"github.com/uanderson/pockee/database"
	"github.com/uanderson/pockee/echox"
	"github.com/uanderson/pockee/errorsx"
	"time"
)

type Service struct {
	dao  *dao.Queries
	pool *pgxpool.Pool
}

func NewService(database *database.Database) *Service {
	return &Service{dao: dao.New(database.Pool), pool: database.Pool}
}

func (s *Service) ExistsContactByID(ctx context.Context, ID string) (bool, error) {
	return s.dao.ExistsContactByID(ctx, dao.ExistsContactByIDParams{
		ID:     ID,
		UserID: echox.GetUserID(ctx),
	})
}

func (s *Service) GetContacts(ctx context.Context) ([]dao.Contact, error) {
	contacts, err := s.dao.GetContacts(ctx, echox.GetUserID(ctx))
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func (s *Service) GetContactByID(ctx context.Context, ID string) (dao.Contact, error) {
	exists, err := s.ExistsContactByID(ctx, ID)
	if err != nil {
		return dao.Contact{}, err
	}

	if !exists {
		return dao.Contact{}, errorsx.ContactNotFound
	}

	return s.dao.GetContactByID(ctx, dao.GetContactByIDParams{
		ID:     ID,
		UserID: echox.GetUserID(ctx),
	})
}

func (s *Service) CreateContact(ctx context.Context, input CreateContactInput) error {
	err := s.dao.CreateContact(ctx, dao.CreateContactParams{
		ID:     autoid.New(),
		Name:   input.Name,
		Email:  input.Email,
		Phone:  input.Phone,
		PixKey: input.PixKey,
		UserID: echox.GetUserID(ctx),
	})

	return err
}

func (s *Service) UpdateContact(ctx context.Context, input UpdateContactInput) error {
	contact, err := s.GetContactByID(ctx, input.ID)
	if err != nil {
		return err
	}

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	qtx := s.dao.WithTx(tx)

	err = qtx.UpdateContact(ctx, dao.UpdateContactParams{
		ID:     input.ID,
		Name:   input.Name,
		Email:  input.Email,
		Phone:  input.Phone,
		PixKey: input.PixKey,
		UserID: echox.GetUserID(ctx),
	})
	if err != nil {
		return err
	}

	err = qtx.CreateContactHistory(ctx, dao.CreateContactHistoryParams{
		ID:     autoid.New(),
		Name:   contact.Name,
		Email:  contact.Email,
		Phone:  contact.Phone,
		PixKey: contact.PixKey,
		EffectiveAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		ContactID: contact.ID,
	})
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (s *Service) DeleteContact(ctx context.Context, input DeleteContactInput) error {
	exists, err := s.ExistsContactByID(ctx, input.ID)
	if err != nil {
		return err
	}

	if !exists {
		return errorsx.ContactNotFound
	}

	return s.dao.SoftDeleteContact(ctx, dao.SoftDeleteContactParams{
		ID: input.ID,
		DeletedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		UserID: echox.GetUserID(ctx),
	})
}
