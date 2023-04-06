package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// ! login / register
type Staff struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Position   string `json:"position"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Experience int    `json:"experience"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginForm struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type RegisterForm struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Attendance struct {
	ID      int    `json:"id"`
	StaffID int    `json:"staff_id"`
	Date    string `json:"date"`
	Time    string `json:"time"`
	XP      int    `json:"xp"`
}

// ! attedence
type AttendanceForm struct {
	StaffID int    `json:"staff_id" form:"staff_id" validate:"required"`
	Date    string `json:"date" form:"date" validate:"required"`
	Time    string `json:"time" form:"time" validate:"required"`
	XP      int    `json:"xp" form:"xp" validate:"required"`
}

// ! staff list
type StaffListResponse struct {
	Status string  `json:"status"`
	Data   []Staff `json:"data"`
}

func main() {
	// Initialize database connection
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_kehadiran")
	if err != nil {
		fmt.Println("Error opening database:", err.Error())
		return
	}
	defer db.Close()

	// Initialize Echo instance
	e := echo.New()

	// Set up the custom validator
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "berhasil data awal")
	})

	// ! Login route
	e.POST("/login", func(c echo.Context) error {
		// Bind and validate login forms data
		form := new(LoginForm)
		if err := c.Bind(form); err != nil {
			return c.JSON(http.StatusBadRequest, Response{"error", "Invalid login form data"})
		}
		if err := c.Validate(form); err != nil {
			return c.JSON(http.StatusBadRequest, Response{"error", "Invalid login form data"})
		}

		// Query user from database
		var user User
		row := db.QueryRow("SELECT id, name, email, password FROM staff WHERE email = ?", form.Email)
		if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			fmt.Printf("Error scanning row: %v\n", err) // Add this line for debugging
			return c.JSON(http.StatusUnauthorized, Response{"error", "Invalid email or password"})
		}
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given email")
			return c.JSON(http.StatusUnauthorized, Response{"error", "Invalid email or password"})
		}
		fmt.Printf("User from DB: %+v\n", user) // Add this line for debugging
		fmt.Printf("Form data: %+v\n", form)    // Add this line for debugging

		// Validate password
		if user.Password != form.Password {
			return c.JSON(http.StatusUnauthorized, Response{"error", "Invalid email or password"})
		}

		// Generate token and return response
		token := generateToken(user.ID)
		return c.JSON(http.StatusOK, Response{"success", token})
	})

	// ! Register route
	e.POST("/register", func(c echo.Context) error {
		// Bind and validate register form data
		form := new(RegisterForm)
		if err := c.Bind(form); err != nil {
			return c.JSON(http.StatusBadRequest, Response{"error", fmt.Sprintf("Invalid register form data: %v", err)})
		}
		if err := c.Validate(form); err != nil {
			return c.JSON(http.StatusBadRequest, Response{"error", fmt.Sprintf("Invalid register form data: %v", err)})
		}

		// Check if user already exists
		var count int
		row := db.QueryRow("SELECT COUNT(*) FROM staff WHERE email = ?", form.Email)
		if err := row.Scan(&count); err != nil {
			return c.JSON(http.StatusInternalServerError, Response{"error", "Internal server error"})
		}
		if count > 0 {
			return c.JSON(http.StatusBadRequest, Response{"error", "Email already exists"})
		}

		// Insert user to database
		result, err := db.Exec("INSERT INTO staff (nama_staff, jabatan, email, password) VALUES (?, ?, ?, ?)", form.Name, "", form.Email, form.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Response{"error", "Internal server error"})
		}

		// Get user ID and generate token
		userID, err := result.LastInsertId()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Response{"error", "Internal server error"})
		}
		token := generateToken(int(userID))

		// Return response
		return c.JSON(http.StatusOK, Response{"success", token})
	})

	// ! Route Home
	e.GET("/home", func(c echo.Context) error {
		// Get staff ID and experience from header
		staffIDStr := c.Request().Header.Get("X-Staff-ID")
		xpStr := c.Request().Header.Get("X-Experience")

		if staffIDStr != "" && xpStr != "" {
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
			row := db.QueryRow("SELECT id_staff, nama_staff, jabatan, email, password, xppoint FROM staff WHERE id_staff = ?", staffID)
			if err := row.Scan(&staff.ID, &staff.Name, &staff.Position, &staff.Email, &staff.Password, &staff.Experience); err != nil {
				return c.JSON(http.StatusInternalServerError, Response{"error", "Internal server error"})
			}
			return c.JSON(http.StatusOK, staff)
		}
		return c.JSON(http.StatusBadRequest, Response{"error", "Invalid request"})
	})

	// ! attendence
	e.POST("/attendance", func(c echo.Context) error {
		// Bind and validate attendance form data
		form := new(AttendanceForm)
		if err := c.Bind(form); err != nil {
			return c.JSON(http.StatusBadRequest, Response{"error", "Invalid attendance form data"})
		}
		if err := c.Validate(form); err != nil {
			return c.JSON(http.StatusBadRequest, Response{"error", "Invalid attendance form data"})
		}

		// Insert attendance record into database
		_, err = db.Exec("INSERT INTO attendance_record (id_staff, tgl_masuk, waktu_masuk, status) VALUES (?, ?, ?, ?)", form.StaffID, form.Date, form.Time, "Present")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Response{"error", "Internal server error"})
		}

		// Update staff XP in database
		_, err = db.Exec("UPDATE staff SET xppoint = xppoint + ? WHERE id_staff = ?", form.XP, form.StaffID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Response{"error", "Internal server error"})
		}

		// Return response
		return c.JSON(http.StatusOK, Response{"success", "Attendance recorded successfully"})
	})

	// ! staff data
	e.GET("/staff", func(c echo.Context) error {
		// Query all staff from database
		rows, err := db.Query("SELECT id_staff, nama_staff, jabatan, email, password, xppoint FROM staff")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Response{"error", "Internal server error"})
		}
		defer rows.Close()

		// Iterate through rows and append to staff list
		staffList := []Staff{}
		for rows.Next() {
			var staff Staff
			err := rows.Scan(&staff.ID, &staff.Name, &staff.Position, &staff.Email, &staff.Password, &staff.Experience)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, Response{"error", "Internal server error"})
			}
			staffList = append(staffList, staff)
		}

		// Check for any errors during row iteration
		if err := rows.Err(); err != nil {
			return c.JSON(http.StatusInternalServerError, Response{"error", "Internal server error"})
		}

		// Return response with staff list
		return c.JSON(http.StatusOK, StaffListResponse{"success", staffList})
	})

	// ? Start server
	e.Logger.Fatal(e.Start(":9000"))
}

func generateToken(userID int) string {
	// Generate token using userID and secret key
	return "some_secret_key_" + strconv.Itoa(userID)
}
