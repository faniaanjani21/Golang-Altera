orm dan code strukture

crud menggunakan eho dan gorm dengan codingan crud menggunakan golang

kepanjangan orm object relational mapping

cara install gorm
https://gorm.io/docs/

https://gorm.io/docs/migration.html

powersheel install
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite

cara menjalanakan nya cukup go run main.go

sebener nya bisa di build dengan 1 file saja yang benernamakan main.go namun di tutorial
mengajarkan cara yang mudah di baca dan membagi nya menjadi beberapa folder
seperti
config
controller
lib
models
routes

dan masing masing mempunyai materi yang berbeda dan juga package yang berbeda

pada materi ini akan menjalankan res api di mana kita bisa mengakses nya dengan postman

tutorial
buka file nya praktikum
buka file main.go
jalankan xampp
isi data nya

func InitDB() {
config := Config{
DB_Username: "root",
DB_Password: "",
DB_Port: "3306",
DB_Host: "localhost",
DB_Name: "crud_go",
}

setelah itu jalankan dengan main.go
setelah itu buka postman
buka url nya http://localhost:8080/users

lakukan dengan metode get, post, put, delete

dimana get berfungsi untuk menampilkan data
post berfungsi untuk menambah data
put berfungsi untuk mengubah data
delete berfungsi untuk menghapus data
