package dependency

import (
	"github.com/erik-sostenes/users-api/internal/apps/backend"
	"github.com/erik-sostenes/users-api/internal/apps/backend/handlers"
	"github.com/erik-sostenes/users-api/internal/apps/backend/handlers/user"
	"github.com/erik-sostenes/users-api/internal/mooc/user/business/services/create"
	"github.com/erik-sostenes/users-api/internal/mooc/user/infrastructure/persistence"
	"github.com/erik-sostenes/users-api/internal/shared/domain/bus/command"
	"github.com/erik-sostenes/users-api/internal/shared/infrastructure/db"
	"github.com/labstack/echo/v4"
)

// NewInjector injects all the dependencies to start the app
func NewInjector() (err error) {
	engine := echo.New()

	collection := db.NewMongoDataBase(db.NewMongoDBConfiguration()).Collection("users")

	commandHandler := &create.CreateUserCommandHandler{
		UserCreator: create.UserCreator{
			UserRepository: persistence.NewUserStore(collection),
		},
	}

	bus := make(command.CommandBus[create.UserCommand])
	if err = bus.Record(create.UserCommand{}, commandHandler); err != nil {
		return
	}

	h := handlers.Handler{
		UserHandler: user.NewUserHandler(&bus),
	}

	return backend.NewServer(h, engine).Start()
}
