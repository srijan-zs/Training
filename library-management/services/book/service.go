package book

import (
	"context"
	"github.com/google/uuid"
	errors2 "github.com/srijan-zs/Training/library-management/errors"
	"github.com/srijan-zs/Training/library-management/models"
	"github.com/srijan-zs/Training/library-management/services"
)

type service struct {
	bookStore services.Book
}

func New(book services.Book) service {
	return service{bookStore: book}
}

func validBook(book *models.Book) error {
	switch {
	case book.Name == "":
		return errors2.InvalidParam{Params: []string{"name"}}
	case book.Author == "":
		return errors2.InvalidParam{Params: []string{"author"}}
	case book.Year > 2022:
		return errors2.InvalidParam{Params: []string{"year"}}
	default:
		return nil
	}
}

func (s service) Create(ctx context.Context, book *models.Book) (*models.Book, error) {
	err := validBook(book)
	if err != nil {
		return nil, err
	}

	id, err := s.bookStore.Create(ctx, book)

	book.ID = id

	output, err := s.bookStore.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (s service) Get(ctx context.Context, id uuid.UUID) (models.Book, error) {
	book, err := s.bookStore.Get(ctx, id)
	if err != nil {
		return models.Book{}, errors2.EntityNotFound{}
	}

	return book, nil
}

func (s service) Update(ctx context.Context, id uuid.UUID, book *models.Book) (*models.Book, error) {
	err := validBook(book)
	if err != nil {
		return nil, err
	}

	resp, err := s.bookStore.Update(ctx, id, book)
	if err != nil {
		return nil, errors2.EntityNotFound{}
	}

	return resp, nil
}

func (s service) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.bookStore.Delete(ctx, id)
	if err != nil {
		return errors2.EntityNotFound{}
	}

	return nil
}
