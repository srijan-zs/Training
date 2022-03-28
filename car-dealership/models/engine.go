package models

import "github.com/google/uuid"

// Engine struct contains all the fields for details about the engine
type Engine struct {
	ID           uuid.UUID `json:"id"`
	Displacement int       `json:"displacement,omitempty"`
	NCylinders   int       `json:"nCylinder,omitempty"`
	Range        int       `json:"extent,omitempty"`
}
