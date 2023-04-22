package user

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/erik-sostenes/users-api/internal/mooc/user/business/services/create"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"sync"
)

func (u *user) Create() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		file, _, err := c.Request().FormFile("users")
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "The .csv file is missing"})
		}
		defer file.Close()

		r := csv.NewReader(file)

		var wg sync.WaitGroup
		for {
			user, err := r.Read()
			if err == io.EOF {
				break
			}

			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{"error": "An error occurred on the server"})
			}

			wg.Add(1)
			go func(user []string) {
				defer wg.Done()

				command := create.UserCommand{
					Id:       user[0],
					Name:     user[1],
					LastName: user[2],
				}

				if err := u.Dispatch(context.Background(), command); err != nil {
					_ = c.JSON(http.StatusCreated, echo.Map{"error": err.Error()})
					return
				}
			}(user)
			_ = c.JSON(http.StatusCreated, echo.Map{"message": fmt.Sprintf("User %s processed", user[0])})
		}

		_ = c.JSON(http.StatusCreated, echo.Map{"message": "users are being processed"})
		wg.Wait()
		return
	}
}
