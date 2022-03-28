package services

import (
	"context"

	"github.com/zopsmart/GoLang-Interns-2022/filters"
	"github.com/zopsmart/GoLang-Interns-2022/models"

	"github.com/google/uuid"
)

type Car interface {
	Create(ctx context.Context, car *models.Car) error
	GetAll(ctx context.Context, filter filters.Car) ([]models.Car, error)
	GetByID(ctx context.Context, id uuid.UUID) (models.Car, error)
	Update(ctx context.Context, id uuid.UUID, car *models.Car) (*models.Car, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type Engine interface {
	Create(ctx context.Context, engine *models.Engine) (uuid.UUID, error)
	GetByID(ctx context.Context, id uuid.UUID) (models.Engine, error)
	Update(ctx context.Context, id uuid.UUID, engine *models.Engine) (*models.Engine, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
