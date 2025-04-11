package model

import "github.com/google/uuid"

type Supplier struct {
	UUID    uuid.UUID
	Name    string
	Contact string
}
