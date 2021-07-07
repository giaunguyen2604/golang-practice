package handler

import (
	"03-middleware/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {

	return c.JSON(http.StatusOK, &models.LoginResponse{
		Token: "123456",
	})
}
