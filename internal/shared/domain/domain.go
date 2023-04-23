package domain

import (
	errs "github.com/erik-sostenes/users-api/internal/shared/domain/errors"
	"strings"
)

// String receives a value to verify if the format is correct
type String string

// The Validate method validates if the value is a string and is not empty, if incorrect returns an errors.StatusUnprocessableEntity
func (s String) Validate() (string, error) {
	if strings.TrimSpace(string(s)) == "" {
		return "", errs.StatusUnprocessableEntity("Value not found")
	}
	return string(s), nil
}
