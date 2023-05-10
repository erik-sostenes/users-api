package persistence

import (
	"context"
	"errors"
	"github.com/erik-sostenes/users-api/internal/mooc/user/business/domain"
	"github.com/erik-sostenes/users-api/internal/shared/infrastructure/db"
	"testing"
)

func TestUserStore_Save(t *testing.T) {
	tsc := map[string]struct {
		user struct {
			id       string
			name     string
			lastName string
		}
		repository    domain.UserRepository[domain.UserId, domain.User]
		expectedError error
	}{
		"Given a valid non-existing user, it will be registered in the mongo collection": {
			user: struct {
				id       string
				name     string
				lastName string
			}{
				"05b15d71-ff47-4695-85d4-3ea3e5c12d60", "Yael", "Castro",
			},
			repository: func() domain.UserRepository[domain.UserId, domain.User] {
				repository := db.NewMongoDataBase(
					db.NewMongoDBConfiguration()).Collection("users")

				return NewUserStore(repository)
			}(),
			expectedError: nil,
		},
		"Given an valid existing user, it will not be registered in the mongo collection": {
			user: struct {
				id       string
				name     string
				lastName string
			}{
				"a37b054e-cb54-443a-953d-261b10476cb3", "Erik", "Sostenes",
			},
			repository: func() domain.UserRepository[domain.UserId, domain.User] {
				repository := db.NewMongoDataBase(
					db.NewMongoDBConfiguration()).Collection("users")

				return NewUserStore(repository)
			}(),
			expectedError: ErrDuplicateUser,
		},
	}

	// SetUp prepare configuration before running integration tests all
	repository := db.NewMongoDataBase(
		db.NewMongoDBConfiguration()).Collection("users")

	store := NewUserStore(repository)

	user, err := domain.NewUser("a37b054e-cb54-443a-953d-261b10476cb3", "Erik", "Sostenes")
	if err != nil {
		t.Fatal(err)
	}

	if err := store.Save(context.TODO(), user.UserId, user); err != nil {
		t.Fatal(err)
	}

	// Teardown reset configuration after running integration tests all
	t.Cleanup(func() {
		if err := store.Delete(context.TODO(), user.UserId); err != nil {
			t.Fatal(err)
		}
	})

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			user, err := domain.NewUser(ts.user.id, ts.user.name, ts.user.lastName)
			if err != nil {
				t.Skip(err)
			}

			t.Cleanup(func() {
				if err := ts.repository.Delete(context.TODO(), user.UserId); err != nil {
					t.Fatal(err)
				}
			})

			err = ts.repository.Save(context.TODO(), user.UserId, user)
			if !errors.Is(err, ts.expectedError) {
				t.Errorf("%v error was expected, but %s error was obtained", ts.expectedError, err)
			}
		})
	}
}
