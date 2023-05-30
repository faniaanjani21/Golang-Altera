package main

import (
	"rest/config"
	"rest/routes"
)

func main() {
	// Inisialisasi koneksi database
	config.InitDB()
	// Menutup koneksi database ketika fungsi main selesai
	defer config.DB.Close()
	// Melakukan migrasi awal untuk struktur tabel
	config.InitialMigration()

	// Membuat instance baru dari Echo dengan rute yang telah ditentukan
	e := routes.New()
	// Memulai server pada port 9000 dan mencatat jika terjadi kesalahan
	e.Logger.Fatal(e.Start(":9000"))
}
