package handler

import (
	errs "github.com/erik-sostenes/users-api/internal/shared/domain/errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ErrorHandler(ctx echo.Context, err error) error {
	switch err.(type) {
	case errs.StatusBadRequest:
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	case errs.StatusUnprocessableEntity:
		return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{"error": err.Error()})
	case errs.StatusNotFound:
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	default:
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
}
