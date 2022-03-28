package brand

import (
	"database/sql/driver"
	"encoding/json"
	"strings"

	errors2 "github.com/zopsmart/GoLang-Interns-2022/errors"
)

type Brand int

// enum type declaration for given brands
const (
	Tesla Brand = iota + 1
	Porsche
	Ferrari
	Mercedes
	BMW
)

// MarshalJSON method converts type Brand to slice of bytes
// nolint:goconst // string used to marshal, unmarshal
func (b Brand) MarshalJSON() ([]byte, error) {
	var s string

	switch b {
	default:
		s = "unknown"
	case Tesla:
		s = "tesla"
	case Porsche:
		s = "porsche"
	case Ferrari:
		s = "ferrari"
	case Mercedes:
		s = "mercedes"
	case BMW:
		s = "bmw"
	}

	return json.Marshal(s)
}

// UnmarshalJSON method converts slice of bytes to type Brand
func (b *Brand) UnmarshalJSON(p []byte) error {
	var s string

	if err := json.Unmarshal(p, &s); err != nil {
		return err
	}

	switch strings.ToLower(s) {
	default:
		*b = 0
	case "tesla":
		*b = Tesla
	case "porsche":
		*b = Porsche
	case "ferrari":
		*b = Ferrari
	case "mercedes":
		*b = Mercedes
	case "bmw":
		*b = BMW
	}

	return nil
}

// IsValid method checks whether the given brand belongs to our collection or not
func (b Brand) IsValid() bool {
	return b >= Tesla && b <= BMW
}

// Value method converts type Brand to value that driver is able to handle
func (b Brand) Value() (driver.Value, error) {
	switch b {
	case Tesla:
		return "tesla", nil
	case Porsche:
		return "porsche", nil
	case Ferrari:
		return "ferrari", nil
	case Mercedes:
		return "mercedes", nil
	case BMW:
		return "bmw", nil
	}

	return nil, errors2.InvalidParam{Params: []string{"brand"}}
}

// Scan method converts driver value to enum type
func (b *Brand) Scan(value interface{}) error {
	brand, _ := value.([]byte)

	switch string(brand) {
	case "tesla":
		*b = Tesla
	case "porsche":
		*b = Porsche
	case "ferrari":
		*b = Ferrari
	case "mercedes":
		*b = Mercedes
	case "bmw":
		*b = BMW
	default:
		return errors2.InvalidParam{Params: []string{"brand"}}
	}

	return nil
}
