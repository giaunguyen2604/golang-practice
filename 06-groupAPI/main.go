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

	isLogedin := middleware.JWT([]byte("mysecretkey"))
	isAdmin := mdw.IsAdminMdw

	server.GET("/", handler.Hello, isLogedin)

	server.GET("/admin", handler.Hello, isLogedin, isAdmin)
	server.POST("/login", handler.Login, middleware.BasicAuth(mdw.BasicAuth))

	groupv2 := server.Group("/v2")
	groupv2.GET("/hello", handler.Hellov2)

	groupUser := server.Group("/user")

	groupUser.GET("/add", handler.AddUser)
	groupUser.GET("/update", handler.UpdateUser)
	groupUser.GET("/delete", handler.DeleteUser)
	server.Logger.Fatal(server.Start(":3000"))

}

type LoginResquest struct {
	Username string `json:"username" form:"username" xml:"username" query:"username"`
	Password string `json:"password" form:"password" xml:"password" query:"password"`
}
