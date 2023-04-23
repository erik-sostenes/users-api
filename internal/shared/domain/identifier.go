package domain

import (
	"fmt"
	errs "github.com/erik-sostenes/users-api/internal/shared/domain/errors"
	"github.com/google/uuid"
)

// Identifier receives a value to verify if the format is correct
type Identifier string

// ParseUuID validate i	f the format the values is a UuID
func (i Identifier) ParseUuID() (string, error) {
	id, err := uuid.Parse(string(i))
	if err != nil {
		return "", errs.StatusUnprocessableEntity(fmt.Sprintf("incorrect %s uuid unique identifier", string(i)))
	}

	return id.String(), nil
}
