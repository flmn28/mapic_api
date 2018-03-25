package main

import (
	"github.com/johskw/mapic_api/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", handler.PostUser)
	e.POST("/login", handler.Login)

	r := e.Group("/restricted")

	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", handler.Restricted)
	r.GET("/locations/:id", handler.GetLocation)
	r.GET("/locations", handler.GetAllLocations)
	r.POST("/locations", handler.PostLocation)
	r.PUT("/locations/:id", handler.PutLocation)
	r.DELETE("/locations/:id", handler.DeleteLocation)

	r.GET("/users/:id", handler.GetUser)
	r.POST("/users", handler.PostUser)
	r.PUT("/users/:id", handler.PutUser)
	r.DELETE("/users/:id", handler.DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}
