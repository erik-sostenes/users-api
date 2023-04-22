package user

import (
	"encoding/csv"
	"github.com/erik-sostenes/users-api/internal/mooc/user/business/services/create"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (u *user) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, _, err := c.Request().FormFile("users")
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "The .csv file is missing"})
		}
		defer file.Close()

		r := csv.NewReader(file)
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{"error": "An error occurred on the server"})
			}
			command := create.UserCommand{
				Id:       record[0],
				Name:     record[1],
				LastName: record[2],
			}

			_ = u.Dispatch(c.Request().Context(), command)
		}

		return c.JSON(http.StatusCreated, echo.Map{"message": "users are being processed"})
	}
}
