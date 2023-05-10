package db

import (
	"errors"
	"testing"
)

func TestNewMongoClient(t *testing.T) {
	t.Run("Given a correct configuration, MongoDB will connect", func(t *testing.T) {
		_, err := NewMongoClient(NewMongoDBConfiguration())

		if !errors.Is(err, nil) {
			t.Errorf("%v error was expected, but %v error was obtained", nil, err)
		}
	})
}
