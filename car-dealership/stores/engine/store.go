package engine

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zopsmart/GoLang-Interns-2022/errors"
	"github.com/zopsmart/GoLang-Interns-2022/models"

	"github.com/google/uuid"
)

type store struct {
	db *sql.DB
}

// New is the factory function for store struct
func New(db *sql.DB) store { // nolint:revive // store is not exported so factory function is the only way to use the methods
	return store{db: db}
}

// Create method creates a new engine and add the data to store
func (s store) Create(ctx context.Context, engine *models.Engine) (uuid.UUID, error) {
	id := uuid.New()

	_, err := s.db.ExecContext(ctx, "INSERT INTO engines (id, displacement, nCylinder, extent) VALUES (?,?,?,?)",
		id.String(), engine.Displacement, engine.NCylinders, engine.Range)

	if err != nil {
		fmt.Println(err)
		return uuid.Nil, errors.DB{Err: err}
	}

	return id, nil
}

// GetByID method extracts the engine of a given id
func (s store) GetByID(ctx context.Context, id uuid.UUID) (models.Engine, error) {
	var engine models.Engine

	resp := s.db.QueryRowContext(ctx, "SELECT * FROM engines WHERE id = ?", id.String())
	err := resp.Scan(&engine.ID, &engine.Displacement, &engine.NCylinders, &engine.Range)

	if err != nil {
		return models.Engine{}, errors.DB{Err: err}
	}

	return engine, nil
}

// Update method modifies the details for a given engine
func (s store) Update(ctx context.Context, id uuid.UUID, engine *models.Engine) (*models.Engine, error) {
	_, err := s.db.ExecContext(ctx, "UPDATE engines SET displacement=?, nCylinder=?, extent=? WHERE id = ?",
		engine.Displacement, engine.NCylinders, engine.Range, id.String())

	if err != nil {
		return nil, errors.DB{Err: err}
	}

	engine.ID = id

	return engine, nil
}

// Delete method deletes the engine of a specific id
func (s store) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM engines WHERE id = ?", id.String())

	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
