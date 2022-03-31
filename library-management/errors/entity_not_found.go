package errors

import (
	"fmt"
	"github.com/google/uuid"
)

type EntityNotFound struct {
	Entity string
	ID     uuid.UUID
}

func (enf EntityNotFound) Error() string {
	return fmt.Sprintf("entity: %v not found, id: %v", enf.Entity, enf.ID)
}
