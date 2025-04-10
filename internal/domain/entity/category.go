package entity

import (
	"errors"
	"github.com/google/uuid"
)

type Category struct {
	UUID        uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}

func (entity *Category) ValidateRules() error {
	if entity.UUID == uuid.Nil {
		return errors.New("invalid uuid")
	}

	if entity.Name == "" || len(entity.Name) > 255 {
		return errors.New("invalid name")
	}

	if entity.Name == "" || len(entity.Description) > 255 {
		return errors.New("invalid description")
	}

	return nil
}
