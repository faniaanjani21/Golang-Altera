package config

import (
	"fmt"

	"rest/models" // Tambahkan impor ini

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// InitDB digunakan untuk menginisialisasi koneksi ke database
func InitDB() {
	// Konfigurasi database Anda
	config := struct {
		DB_Username string
		DB_Password string
		DB_Port     string
		DB_Host     string
		DB_Name     string
	}{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_go",
	}

	// Membuat string koneksi menggunakan konfigurasi yang telah didefinisikan
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	// Membuka koneksi ke database menggunakan string koneksi
	DB, err = gorm.Open("mysql", connectionString)
	// Jika ada kesalahan saat membuka koneksi, hentikan eksekusi program
	if err != nil {
		panic(err)
	}
}

// InitialMigration digunakan untuk menjalankan migrasi awal pada database
func InitialMigration() {
	// Migrasi model User
	DB.AutoMigrate(&models.User{}) // Ganti menjadi models.User{}
}
