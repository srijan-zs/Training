package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/srijan-zs/Training/library-management/models"
)

type Book interface {
	Create(ctx context.Context, book *models.Book) (uuid.UUID, error)
	Get(ctx context.Context, id uuid.UUID) (models.Book, error)
	Update(ctx context.Context, id uuid.UUID, book *models.Book) (*models.Book, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
