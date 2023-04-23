package persistence

import (
	"context"
	"fmt"
	"github.com/erik-sostenes/users-api/internal/mooc/user/business/domain"
	errs "github.com/erik-sostenes/users-api/internal/shared/domain/errors"
	"sync"
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
