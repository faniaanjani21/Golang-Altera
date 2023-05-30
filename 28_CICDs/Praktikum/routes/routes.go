package routes

import (
	"rest/controllers"

	"github.com/labstack/echo"
)

// New mengembalikan instance baru dari Echo dengan rute yang telah ditentukan
func New() *echo.Echo {
	e := echo.New()

	// Rute untuk mengambil semua pengguna
	e.GET("/users", controllers.GetUsersController)

	return e
}
