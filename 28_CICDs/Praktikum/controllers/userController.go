package controllers

import (
	"net/http"

	"rest/lib/database"

	"github.com/labstack/echo"
)

// GetUsersController digunakan untuk mengambil semua pengguna dari database
func GetUsersController(c echo.Context) error {
	// Mengambil semua pengguna dari database menggunakan fungsi GetUsers()
	users, err := database.GetUsers()
	// Jika terjadi kesalahan saat mengambil pengguna, kirim respons dengan status BadRequest dan pesan kesalahan
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error getting users",
		})
	}
	// Jika berhasil mengambil pengguna, kirim respons dengan status OK dan daftar pengguna
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}
