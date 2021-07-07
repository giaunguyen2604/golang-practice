package mdw

import "github.com/labstack/echo/v4"

func BasicAuth(username string, password string, c echo.Context) (bool, error) {
	if username != "Admin" || password != "password" {
		return false, nil
	}
	return true, nil
}
