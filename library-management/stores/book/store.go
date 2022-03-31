package book

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	errors2 "github.com/srijan-zs/Training/library-management/errors"
	"github.com/srijan-zs/Training/library-management/models"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) store {
	return store{db: db}
}

func (s store) Create(ctx context.Context, book *models.Book) (uuid.UUID, error) {
	id := uuid.New()

	_, err := s.db.ExecContext(ctx, "INSERT INTO books (id, name, year, author) VALUES (?, ?, ?, ?)",
		id.String(), book.Name, book.Year, book.Author)

	if err != nil {
		fmt.Println(err)
		return uuid.Nil, errors2.DB{Err: err}
	}

	return id, nil
}

func (s store) Get(ctx context.Context, id uuid.UUID) (models.Book, error) {
	var book models.Book

	resp := s.db.QueryRowContext(ctx, "SELECT * FROM books WHERE id = ?", id.String())
	err := resp.Scan(&book.ID, &book.Name, &book.Year, &book.Author)

	switch {
	case err == sql.ErrNoRows:
		return models.Book{}, errors2.EntityNotFound{Entity: "book", ID: id}
	case err != nil:
		return models.Book{}, errors2.DB{Err: err}
	default:
		return book, nil
	}
}

func (s store) Update(ctx context.Context, id uuid.UUID, book *models.Book) (*models.Book, error) {
	_, err := s.db.ExecContext(ctx, "UPDATE books SET name=?, year=?, author=? WHERE id=?",
		book.Name, book.Year, book.Author, id.String())

	if err != nil {
		return nil, errors2.DB{Err: err}
	}

	return book, nil
}

func (s store) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM books WHERE id = ?", id.String())

	if err != nil {
		return errors2.DB{Err: err}
	}

	return nil
}
