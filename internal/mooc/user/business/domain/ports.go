package domain

import "context"

type UserRepository[K comparable, V any] interface {
	Save(ctx context.Context, k K, v V) error
}
