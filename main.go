package main

import (
	"github.com/flmn28/mapic_api/handler"
	"github.com/flmn28/mapic_api/models/location"
	"github.com/flmn28/mapic_api/models/user"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	db, err := gorm.Open("mysql", "root@/mapic_api_development?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&location.Location{}, &user.User{})
}

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/locations/:id", handler.GetLocation)
	e.GET("/locations", handler.GetAllLocations)
	e.POST("/locations", handler.PostLocation)
	e.PUT("/locations/:id", handler.PutLocation)
	e.DELETE("/locations/:id", handler.DeleteLocation)

	e.GET("/users/:id", handler.GetUser)
	e.POST("/users", handler.PostUser)
	e.PUT("/users/:id", handler.PutUser)
	e.DELETE("/users/:id", handler.DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}
