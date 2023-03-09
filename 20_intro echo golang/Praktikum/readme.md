jalankan di terminal

go mod init rest
https://echo.labstack.com/guide/#installation
go get -u github.com/labstack/o/v4

penjelasan
routing merupakan sebuah jalur api
contoh
localhost:8080/user
akan mengakses controller dengan fungsi get user

+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
contoh nya
package main

import (
"encoding/json"
"net/http"

    "github.com/labstack/echo"

)

type User struct {
Email string
Name string
}

func main() {
e := echo.New()

    // routing
    e.GET("/hello", helloController)

    e.Start(":8000")

}

// response "helloworld"
func helloController(c echo.Context) error {

    user := User{Email: "academy@altera.id", Name: "altera academy"}

    // marshal data ke bentuk json
    response, err := json.Marshal(user)
    if err != nil {
    	return c.String(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, string(response))

}

ntr di postman method get
http://localhost:8000/hello

+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

Untuk menjalankan API yang telah di-code-kan dengan Echo, dapat dilakukan langkah-langkah berikut:

Jalankan program Go (pastikan sudah terinstall) dengan mengetikkan go run main.go pada terminal dari directory project.

Buka Postman.

Untuk melakukan GET request ke endpoint '/users', pilih method GET dan masukkan URL http://localhost:8000/users. Kemudian, klik tombol Send.

get_all_users

Untuk melakukan POST request ke endpoint '/users' untuk membuat user baru, pilih method POST dan masukkan URL http://localhost:8000/users. Pada tab Body, pilih raw dan content type JSON. Isi field sesuai atribut untuk membuat user baru. Klik tombol Send.
create_user

Hasil post request akan tampil pada Response body di bawahnya.
create_user_response
Catatan: Pastikan method dan URL yang digunakan pada Postman sama dengan yang digunakan pada kode program Echo.
