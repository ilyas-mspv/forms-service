package postgres

import (
	"context"
	"fmt"
	"forms/internal/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sqlx.DB
}

func New(connectionUrl string) (*Storage, error) {
	const op = "storage.postres.New"
	db, err := sqlx.Connect("postgres", connectionUrl)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Storage{db: db}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) Ping() error {
	return s.db.Ping()
}

func (s *Storage) GetAll(ctx context.Context) ([]*models.Form, error) {
	const op = "storage.postgres.GetAll"

	var forms []*models.Form
	err := s.db.SelectContext(ctx, &forms, "SELECT * FROM form")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return forms, nil
}

func (s *Storage) CreateForm(ctx context.Context, form *models.Form) (int64, error) {
	const op = "storage.postgres.StoreForm"
	res, err := s.db.ExecContext(ctx, "INSERT INTO form (name, identifier, description) VALUES ($1, $2, $3)", form.Name, form.Identifier, form.Description)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (s *Storage) Form(ctx context.Context, identifier string) (*models.Form, error) {
	const op = "storage.postgres.Form"
	form := &models.Form{}
	err := s.db.GetContext(ctx, form, "SELECT * FROM form WHERE identifier = $1", identifier)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return form, nil
}
