package routes

import (
	"Praktikum/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", controllers.Home)
	e.GET("/users", controllers.GetUsers)
	e.GET("/users/:id", controllers.GetUser)
	e.POST("/users", controllers.CreateUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)

	return e
}
