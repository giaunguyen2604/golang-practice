package handler

import (
	"03-middleware/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	x := &models.X{
		Text: "Hello World",
	}
	return c.JSON(http.StatusOK, x)
}
