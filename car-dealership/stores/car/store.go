package car

import (
	"context"
	"database/sql"
	"fmt"

	errors2 "github.com/zopsmart/GoLang-Interns-2022/errors"
	"github.com/zopsmart/GoLang-Interns-2022/filters"
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

// Create method creates a new car and add the data to store
func (s store) Create(ctx context.Context, car *models.Car) error {
	_, err := s.db.ExecContext(ctx, "INSERT INTO cars (id, name, year, brand, fuel, engine_id) VALUES (?,?,?,?,?,?)",
		car.ID.String(), car.Name, car.Year, car.Brand, car.Fuel, car.Engine.ID.String())

	if err != nil {
		fmt.Println(err)
		return errors2.DB{Err: err}
	}

	return nil
}

// GetAll method extracts cars using a filter for brand and engine
func (s store) GetAll(ctx context.Context, filter filters.Car) ([]models.Car, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT * FROM cars WHERE brand = ?", filter.Brand)

	if err != nil {
		return nil, errors2.DB{Err: err}
	}

	defer rows.Close()

	cars := make([]models.Car, 0)

	for rows.Next() {
		var car models.Car

		err = rows.Scan(&car.ID, &car.Name, &car.Year, &car.Brand, &car.Fuel, &car.Engine.ID)
		if err != nil {
			return nil, errors2.DB{Err: err}
		}

		cars = append(cars, car)
	}

	if err = rows.Err(); err != nil {
		return nil, errors2.DB{Err: err}
	}

	return cars, nil
}

// GetByID method extracts the car of a given id
func (s store) GetByID(ctx context.Context, id uuid.UUID) (models.Car, error) {
	var car models.Car

	resp := s.db.QueryRowContext(ctx, "SELECT * FROM cars WHERE id = ?", id.String())
	err := resp.Scan(&car.ID, &car.Name, &car.Year, &car.Brand, &car.Fuel, &car.Engine.ID)

	switch {
	case err == sql.ErrNoRows:
		return models.Car{}, errors2.EntityNotFound{Entity: "car", ID: id}
	case err != nil:
		return models.Car{}, errors2.DB{Err: err}
	default:
		return car, nil
	}
}

// Update method modifies the details for a given car
func (s store) Update(ctx context.Context, id uuid.UUID, car *models.Car) (*models.Car, error) {
	_, err := s.db.ExecContext(ctx, "UPDATE cars SET name=?, year=?, brand=?, fuel=?, engine_id=? WHERE id=?",
		car.Name, car.Year, car.Brand, car.Fuel, id.String(), id.String())

	if err != nil {
		return nil, errors2.DB{Err: err}
	}

	return car, nil
}

// Delete method deletes the car of a specific id
func (s store) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM cars WHERE id = ?", id.String())

	if err != nil {
		return errors2.DB{Err: err}
	}

	return nil
}
