package models

import (
	"github.com/zopsmart/GoLang-Interns-2022/types/brand"
	"github.com/zopsmart/GoLang-Interns-2022/types/fuel"

	"github.com/google/uuid"
)

// Car struct contains all the fields for details about the car
type Car struct {
	ID     uuid.UUID   `json:"id"`
	Name   string      `json:"name"`
	Year   int         `json:"year"`
	Brand  brand.Brand `json:"brand"`
	Fuel   fuel.Fuel   `json:"fuel"`
	Engine Engine      `json:"engine"`
}
