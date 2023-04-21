package user

import (
	"encoding/csv"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
)

type UserRequest struct {
	ID       string
	Name     string
	LastName string
}

func (u *user) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, _, err := c.Request().FormFile("users")
		if err != nil {
			return c.JSON(http.StatusBadRequest, "The .csv file is missing")
		}
		defer file.Close()

		r := csv.NewReader(file)
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return c.String(http.StatusInternalServerError, "An error occurred on the server")
			}
			log.Println(record)
		}

		return c.JSON(http.StatusCreated, "users created")
	}
}
