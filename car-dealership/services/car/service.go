package car

import (
	"context"

	errors2 "github.com/zopsmart/GoLang-Interns-2022/errors"
	"github.com/zopsmart/GoLang-Interns-2022/filters"
	"github.com/zopsmart/GoLang-Interns-2022/models"
	"github.com/zopsmart/GoLang-Interns-2022/services"

	"github.com/google/uuid"
)

type service struct {
	carStore    services.Car
	engineStore services.Engine
}

func New(car services.Car, engine services.Engine) service { // nolint:revive // factory function
	return service{carStore: car, engineStore: engine}
}

// validCar function checks for the car if that passes the given criteria or fails
func validCar(car *models.Car) error {
	switch {
	case car.Name == "":
		return errors2.InvalidParam{Params: []string{"name"}}
	case car.Year < 1980 || car.Year > 2022:
		return errors2.InvalidParam{Params: []string{"year"}}
	case !car.Brand.IsValid():
		return errors2.InvalidParam{Params: []string{"brand"}}
	case !car.Fuel.IsValid():
		return errors2.InvalidParam{Params: []string{"fuel"}}
	default:
		return nil
	}
}

// validEngine function checks for the engine if that passes the given criteria or not
func validEngine(engine models.Engine) error {
	equal := func(x int) bool {
		return x == 0
	}

	greater := func(x int) bool {
		return x > 0
	}

	if validate(engine, equal) || validate(engine, greater) {
		return errors2.InvalidParam{Params: []string{"displacement", "nCylinder", "range"}}
	}

	params := make([]string, 0)

	switch {
	case engine.Displacement < 0:
		params = append(params, "displacement")
	case engine.NCylinders < 0:
		params = append(params, "nCylinder")
	case engine.Range < 0:
		params = append(params, "range")
	}

	if len(params) > 0 {
		return errors2.InvalidParam{Params: params}
	}

	return nil
}

// validate function checks for the valid values of engine properties
func validate(e models.Engine, compare func(int) bool) bool {
	return compare(e.Displacement) && compare(e.NCylinders) && compare(e.Range)
}

// Create method validates the data and sends the data to store layer to create a new car
func (s service) Create(ctx context.Context, car *models.Car) (*models.Car, error) {
	err := validCar(car)
	if err != nil {
		return nil, err
	}

	err = validEngine(car.Engine)
	if err != nil {
		return nil, err
	}

	id, err := s.engineStore.Create(ctx, &car.Engine)
	if err != nil {
		return nil, err
	}

	car.ID = id
	car.Engine.ID = id

	err = s.carStore.Create(ctx, car)
	if err != nil {
		return nil, err
	}

	output, err := s.carStore.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp, err := s.engineStore.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	output.Engine = resp
	output.ID = id

	return &output, nil
}

// GetAll method uses brand and engine as a filter to extract cars with specified details
func (s service) GetAll(ctx context.Context, filter filters.Car) ([]models.Car, error) {
	cars, err := s.carStore.GetAll(ctx, filter)
	if err != nil {
		return nil, err
	}

	if filter.IncludeEngine {
		for i := range cars {
			engine, err := s.engineStore.GetByID(ctx, cars[i].ID)

			if err != nil {
				return cars, err
			}

			cars[i].Engine = engine
		}
	}

	return cars, nil
}

// GetByID method extracts the car of a given id
func (s service) GetByID(ctx context.Context, id uuid.UUID) (models.Car, error) {
	car, err := s.carStore.GetByID(ctx, id)
	if err != nil {
		return models.Car{}, errors2.EntityNotFound{}
	}

	engine, err := s.engineStore.GetByID(ctx, id)
	if err != nil {
		return models.Car{}, errors2.EntityNotFound{}
	}

	car.Engine = engine

	return car, nil
}

// Update method validates the data and sends the data to store layer to modify the details of car
func (s service) Update(ctx context.Context, id uuid.UUID, car *models.Car) (*models.Car, error) {
	err := validCar(car)
	if err != nil {
		return nil, err
	}

	_, err = s.engineStore.Update(ctx, id, &car.Engine)
	if err != nil {
		return nil, errors2.EntityNotFound{}
	}

	resp, err := s.carStore.Update(ctx, id, car)
	if err != nil {
		return nil, errors2.EntityNotFound{}
	}

	return resp, nil
}

// Delete method deleted the car of a given id
func (s service) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.engineStore.Delete(ctx, id)
	if err != nil {
		return errors2.EntityNotFound{}
	}

	err = s.carStore.Delete(ctx, id)
	if err != nil {
		return errors2.EntityNotFound{}
	}

	return nil
}
