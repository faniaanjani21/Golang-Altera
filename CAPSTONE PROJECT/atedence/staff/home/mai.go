package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Staff struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Position   string `json:"position"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Experience int    `json:"experience"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	// Initialize database connection
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/db_kehadiran")
	if err != nil {
		fmt.Println("Error opening database:", err.Error())
		return
	}
	defer db.Close()

	// Initialize Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		// Get staff ID and experience from header
		staffIDStr := c.Request().Header.Get("X-Staff-ID")
		xpStr := c.Request().Header.Get("X-Experience")

		// Convert staff ID and experience to integer
		staffID, err := strconv.Atoi(staffIDStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, Response{"error", "Invalid staff ID"})
		}
		xp, err := strconv.Atoi(xpStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, Response{"error", "Invalid experience value"})
		}

		// Update staff experience in database
		_, err = db.Exec("UPDATE staff SET xppoint = xppoint + ? WHERE id_staff = ?", xp, staffID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Response{"error", "Internal server error"})
		}

		// Query staff from database and return response
		var staff Staff
		row := db.QueryRow("SELECT id_staff, xppoint FROM staff WHERE id_staff = ?", staffID)
		if err := row.Scan(&staff.ID, &staff.Experience); err != nil {
			return c.JSON(http.StatusInternalServerError, Response{"error", "Internal server error"})
		}
		return c.JSON(http.StatusOK, staff)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Dalam kode program di atas, terdapat satu model yaitu Staff yang merepresentasikan struktur tabel yang telah Anda buat di dalam database. Pada bagian utama kode program, terdapat fungsi main() yang menginisialisasi koneksi ke database, inisialisasi instance Echo, menambahkan middleware, dan menambahkan route untuk halaman home.

// Pada route untuk halaman home, kita melakukan query ke database untuk mengambil informasi id_staff dan xppoint dari header yang dikirim oleh client. Setelah itu, kita mengupdate nilai xppoint di dalam database sesuai dengan nilai yang dikirim oleh client. Terakhir, kita mengambil kembali informasi id_staff dan xppoint terbaru dari database dan mengirimkannya sebagai response ke client.
