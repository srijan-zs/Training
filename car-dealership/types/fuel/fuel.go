package fuel

import (
	"database/sql/driver"
	"encoding/json"
	"strings"

	errors2 "github.com/zopsmart/GoLang-Interns-2022/errors"
)

type Fuel int

// enum type declaration for given fuels
const (
	Electric Fuel = iota + 1
	Diesel
	Petrol
)

// MarshalJSON method converts type Fuel to slice of bytes
// nolint:goconst // string used to marshal, unmarshal
func (f Fuel) MarshalJSON() ([]byte, error) {
	var s string

	switch f {
	default:
		s = "unknown"
	case Electric:
		s = "electric"
	case Diesel:
		s = "diesel"
	case Petrol:
		s = "petrol"
	}

	return json.Marshal(s)
}

// UnmarshalJSON method converts slice of bytes to type Fuel
func (f *Fuel) UnmarshalJSON(p []byte) error {
	var s string

	if err := json.Unmarshal(p, &s); err != nil {
		return err
	}

	switch strings.ToLower(s) {
	default:
		*f = 0
	case "electric":
		*f = Electric
	case "diesel":
		*f = Diesel
	case "petrol":
		*f = Petrol
	}

	return nil
}

// IsValid method checks whether the given fuel belongs to our collection or not
func (f Fuel) IsValid() bool {
	return f >= Electric && f <= Petrol
}

// Value method converts type Fuel to value that driver is able to handle
func (f Fuel) Value() (driver.Value, error) {
	switch f {
	case Diesel:
		return "Diesel", nil
	case Petrol:
		return "Petrol", nil
	case Electric:
		return "Electric", nil
	}

	return nil, errors2.InvalidParam{Params: []string{"fuel"}}
}

// Scan method converts driver value to enum type
func (f *Fuel) Scan(value interface{}) error {
	fuel, _ := value.([]byte)

	switch strings.ToLower(string(fuel)) {
	case "diesel":
		*f = Diesel
	case "petrol":
		*f = Petrol
	case "electric":
		*f = Electric
	default:
		return errors2.InvalidParam{Params: []string{"fuel"}}
	}

	return nil
}
