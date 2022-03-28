package errors

import "fmt"

// EntityAlreadyExists struct is defined for duplication errors
type EntityAlreadyExists struct {
	Entity string
}

// Error method specifies the format of the error returned according to its types
func (e EntityAlreadyExists) Error() string {
	return fmt.Sprintf("entity  %v already exists", e.Entity)
}
