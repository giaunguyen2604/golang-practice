package main

import (
	"03-middleware/handler"
	"03-middleware/mdw"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	server := echo.New()

	server.Use(middleware.Logger())

	server.GET("/", handler.Hello)

	server.POST("/login", handler.Login, middleware.BasicAuth(mdw.BasicAuth))

	server.Logger.Fatal(server.Start(":3000"))

}

type LoginResquest struct {
	Username string `json:"username" form:"username" xml:"username" query:"username"`
	Password string `json:"password" form:"password" xml:"password" query:"password"`
}
