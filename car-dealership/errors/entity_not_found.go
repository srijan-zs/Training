package errors

import (
	"fmt"

	"github.com/google/uuid"
)

// EntityNotFound struct is defined for the non-existing entity errors
type EntityNotFound struct {
	Entity string
	ID     uuid.UUID
}

// Error method specifies the format of the error returned according to its types
func (enf EntityNotFound) Error() string {
	return fmt.Sprintf("entity: %v not found, id: %v", enf.Entity, enf.ID)
}
