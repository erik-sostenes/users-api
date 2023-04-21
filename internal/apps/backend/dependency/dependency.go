package dependency

import (
	"github.com/erik-sostenes/users-api/internal/apps/backend"
	"github.com/labstack/echo/v4"
)

// NewInjector injects all the dependencies to start the app
func NewInjector() (err error) {
	engine := echo.New()

	return backend.NewServer(engine).Start()
}
