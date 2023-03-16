# ORM dan Struktur Kode

## ORM (Object-Relational Mapping)

ORM adalah suatu teknologi yang memungkinkan kita untuk mengakses database relasional dengan menggunakan objek-objek yang terdefinisi dalam suatu bahasa pemrograman. Dengan menggunakan ORM, kita dapat memanipulasi data pada database dengan lebih mudah karena kita tidak perlu menulis query SQL secara langsung.

## GORM

GORM adalah salah satu ORM untuk Golang yang cukup populer dan banyak digunakan. GORM menyediakan fitur-fitur yang lengkap, seperti dukungan untuk berbagai macam database, relasi antar tabel, validasi data, dan lain-lain.

Untuk menginstall GORM, kita dapat mengikuti langkah-langkah berikut:

1. Buka terminal atau PowerShell.
2. Jalankan perintah `go get -u gorm.io/gorm`.
3. Jalankan perintah `go get -u gorm.io/driver/sqlite` jika kamu ingin menggunakan database SQLite.

## Struktur Kode

Untuk membuat kode kita lebih terstruktur, kita dapat membagi kode menjadi beberapa direktori dan file. Berikutadalah struktur kode yang dapat digunakan untuk membuat CRUD dengan GORM:

- `config`: berisi konfigurasi untuk koneksi ke database.
- `controller`: berisi fungsi-fungsi yang akan dipanggil saat ada request dari client.
- `lib`: berisi helper function atau library yang dapat digunakan di seluruh aplikasi.
- `models`: berisi definisi model atau struktur tabel pada database.
- `routes`: berisi definisi route atau endpoint yang dapat diakses oleh client.
- `main.go`: file utama yang akan memanggil fungsi-fungsi dari package lain dan menjalankan server.

Dengan membagi kode menjadi beberapa package seperti ini, kita dapat memudahkan proses development dan memperbaiki kode jika terjadi kesalahan pada bagian tertentu.

Untuk menjalankan aplikasi, kita dapat mengetikkan perintah `go run main.go` di terminal atau PowerShell. Setelah itu, kita dapat mengakses aplikasi melalui browser atau aplikasi sejenis Postman dengan memasukkan URL endpoint yang telah kita definisikan di dalam `routes`.
