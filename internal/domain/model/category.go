package model

import "github.com/google/uuid"

type Category struct {
	UUID        uuid.UUID
	Name        string
	Description string
}
