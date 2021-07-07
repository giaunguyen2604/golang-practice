package handler

import (
	"03-middleware/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloAdmin(c echo.Context) error {

	message := "Hello Admin"

	x := &models.X{
		Text: message,
	}
	return c.JSON(http.StatusOK, x)
}
