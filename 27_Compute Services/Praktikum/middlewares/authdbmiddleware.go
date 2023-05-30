package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	// Rute
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "berhasil mendapatkan pengguna berdasarkan id",
			"id":      id,
		})
	})

	e.POST("/users", func(c echo.Context) error {
		// Menangani pembuatan pengguna
		return c.String(http.StatusOK, "Pengguna dibuat")
	})

	// Mulai server
	log.Fatal(e.Start(":8080"))
}
