package handler

import (
	"03-middleware/models"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	username := c.Get("username").(string)
	log.Printf("login with username %s\n", username)
	admin := c.Get("admin").(bool)
	log.Printf("login with admin %v\n", admin)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("mysecretkey"))

	if err != nil {
		log.Printf("signed token err %v\n", err)
		return err
	}

	return c.JSON(http.StatusOK, &models.LoginResponse{
		Token: t,
	})
}
