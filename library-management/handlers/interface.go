package handlers

import (
	"context"
	"github.com/google/uuid"
	"github.com/srijan-zs/Training/library-management/models"
)

type Service interface {
	Create(ctx context.Context, book *models.Book) (*models.Book, error)
	Get(ctx context.Context, id uuid.UUID) (models.Book, error)
	Update(ctx context.Context, id uuid.UUID, book *models.Book) (*models.Book, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
