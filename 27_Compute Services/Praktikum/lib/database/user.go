package database

import (
	"rest/config"
	"rest/models"
)

// GetUsers mengambil semua pengguna dari database
func GetUsers() ([]models.User, error) {
	var users []models.User
	// Mencari semua pengguna dalam database menggunakan objek DB dari paket config
	err := config.DB.Find(&users).Error
	// Mengembalikan daftar pengguna dan kesalahan (jika ada)
	return users, err
}
