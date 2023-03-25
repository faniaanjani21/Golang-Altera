package route

import (
	"prio2/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// buku
	e.GET("/buku", controller.GetBuku)
	e.GET("/buku/:id", controller.GetBukuById)
	e.POST("/buku", controller.CreateBuku)
	e.DELETE("/buku/:id", controller.DeleteBuku)
	e.PUT("/buku/:id", controller.UpdateBuku)

	// user
	e.GET("/user", controller.GetUsers)
	e.GET("/user/:id", controller.GetUserById)
	e.POST("/user", controller.CreateUser)
	e.DELETE("/user/:id", controller.DeleteUser)
	e.PUT("/user/:id", controller.UpdateUser)
	return e

}
