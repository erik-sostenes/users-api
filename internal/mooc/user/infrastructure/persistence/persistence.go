package persistence

import (
	"context"
	"errors"
	"fmt"
	"github.com/erik-sostenes/users-api/internal/mooc/user/business/domain"
	errs "github.com/erik-sostenes/users-api/internal/shared/domain/errors"
	"sync"
)

var (
	ErrDuplicateUser = errors.New("duplicate user with id")
)

type mockUserRepository[K comparable, V any] struct {
	cache map[K]V
	mux   sync.Mutex
}

func NewMockUserRepository[K comparable, V any]() domain.UserRepository[K, V] {
	return &mockUserRepository[K, V]{
		cache: make(map[K]V),
	}
}

// Save saves a resource in a map
// if the resource already exist, returns a StatusBadRequest type error
func (m *mockUserRepository[K, V]) Save(_ context.Context, k K, v V) (err error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	_, ok := m.cache[k]
	if ok {
		err = errs.StatusBadRequest(fmt.Sprintf("resource with id %v already existing", k))
		return
	}

	m.cache[k] = v

	return
}

// Delete removes the resources by an identifier of a map
// if the resource is not found, returns a Not Found type error
func (m *mockUserRepository[K, V]) Delete(_ context.Context, k K) (err error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	_, ok := m.cache[k]

	if !ok {
		err = errs.StatusNotFound(fmt.Sprintf("resource with id %v not found", k))
		return
	}

	delete(m.cache, k)

	return
}

// Find searches a resource by id from a map
// if the resource is not found, returns a Not Found type error
func (m *mockUserRepository[K, V]) Find(_ context.Context, k K) (v V, err error) {
	m.mux.Lock()
	defer m.mux.Unlock()

	v, ok := m.cache[k]
	if !ok {
		err = errs.StatusNotFound(fmt.Sprintf("resource with id %v not found", k))
		return
	}
	return
}
