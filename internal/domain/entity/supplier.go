package entity

import (
	"errors"
	"github.com/google/uuid"
)

type Supplier struct {
	UUID    uuid.UUID `db:"uuid"`
	Name    string    `db:"name"`
	Contact string    `db:"contact"`
}

func (entity *Supplier) ValidateRules() error {
	if entity.UUID == uuid.Nil {
		return errors.New("invalid uuid")
	}

	if entity.Name == "" || len(entity.Name) > 255 {
		return errors.New("invalid name")
	}

	if entity.Contact == "" || len(entity.Contact) > 255 {
		return errors.New("invalid contact")
	}

	return nil
}
