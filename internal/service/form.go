package service

import (
	"context"
	"errors"
	"forms/internal/models"
	"forms/internal/storage"
)

type FormRepository interface {
	GetAll(ctx context.Context) ([]*models.Form, error)
	CreateForm(ctx context.Context, form *models.Form) (int64, error)
	Form(ctx context.Context, identifier string) (*models.Form, error)
}

type FormUseCase struct {
	repository FormRepository
}

func New(repository FormRepository) *FormUseCase {
	return &FormUseCase{repository: repository}
}

func (u *FormUseCase) GetAllForms(ctx context.Context) ([]*models.Form, error) {
	return u.repository.GetAll(ctx)
}

func (u *FormUseCase) CreateForm(ctx context.Context, form *models.Form) (int64, error) {
	_, err := u.repository.Form(ctx, form.Identifier)
	if errors.Is(err, storage.ErrNotFound) {
		return u.repository.CreateForm(ctx, form)
	}
	return 0, storage.ErrAlreadyExists
}

func (u *FormUseCase) Form(ctx context.Context, identifier string) (*models.Form, error) {
	return u.repository.Form(ctx, identifier)
}
