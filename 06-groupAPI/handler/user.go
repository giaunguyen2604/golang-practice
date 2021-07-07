package handler

import (
	"03-middleware/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddUser(c echo.Context) error {

	message := "Add user"

	x := &models.X{
		Text: message,
	}
	return c.JSON(http.StatusOK, x)
}

func UpdateUser(c echo.Context) error {

	message := "Update user"

	x := &models.X{
		Text: message,
	}
	return c.JSON(http.StatusOK, x)
}

func DeleteUser(c echo.Context) error {

	message := "Delete user"

	x := &models.X{
		Text: message,
	}
	return c.JSON(http.StatusOK, x)
}
