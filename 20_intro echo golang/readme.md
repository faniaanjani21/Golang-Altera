Materi – Intro Echo Golangs

Understand the concept of API
What is API?
REST
JSON
HTTP Response Code
Open API
Introduction to POSTMan
Package net/http

API atau Application Programming Interface adalah kumpulan aturan, protokol, dan alat yang digunakan untuk membangun perangkat lunak dan aplikasi. API merupaka interface yang berfungsi menghubungkan program dengan sumber daya atau layanan lainnya yang ada pada sistem operasi atau di dalam aplikasi itu sendiri.

Secara singkat, API adalah cara untuk menjembatani interaksi antara dua program. Dengan menggunakan API, pengembang bisa membuat aplikasi yang dapat berinteraksi dengan aplikasi dan layanan lain secara mudah dan efisien.

API biasanya memiliki beberapa endpoint atau titik akhir yang dapat diakses oleh pengembang dengan melakukan permintaan HTTP yang sesuai dengan aturan-aturan yang telah ditetapkan oleh API tersebut. Setelah permintaan diterima oleh API server, data yang diminta akan dikirim kembali ke aplikasi pengguna dalam format yang telah ditentukan sebelumnya.

Contohnya, Saat Anda menggunakan aplikasi pemesanan tiket pesawat, aplikasi tersebut berkomunikasi dengan API provider tiket pesawat untuk mendapatkan informasi harga tiket, jadwal penerbangan, dan lain-lain. Kemudian aplikasi akan menampilkan hasilnya dengan tampilan yang sudah dibuat sebelumnya.

sepulsa, midtrans,bank

format dalam request dan response
JSON
XML
soap

contoh object notation
{
"name": "John",
"age": 30,
"cars": [
{ "name":"Ford", "models":[ "Fiesta", "Focus", "Mustang" ] },
{ "name":"BMW", "models":[ "320", "X3", "X5" ] },
{ "name":"Fiat", "models":[ "500", "Panda" ] }
]
}

http response code
200 OK
201 Created
202 Accepted
204 No Content
301 Moved Permanently
302 Found
304 Not Modified
400 Bad Request
401 Unauthorized
403 Forbidden
404 Not Found
405 Method Not Allowed
500 Internal Server Error
502 Bad Gateway
503 Service Unavailable

tools yang di gunakan
postman
// membuat text bagus di readme.md
[ ] -postman
[ ] -katalon
[ ] -apache jmeter
[ ] -soapui

SWAPI (Star Wars API) adalah sebuah layanan web yang menyediakan data terkait Star Wars. API ini dapat diakses oleh developer untuk mengintegrasikan data-data tersebut dengan aplikasi atau website mereka.

Beberapa contoh penggunaan SWAPI antara lain:

Membuat website dengan tema Star Wars yang menampilkan detail karakter, planet, dan film dari saga Star Wars.
Membuat aplikasi mobile yang dapat memberikan informasi lengkap tentang suatu karakter atau planet dari Star Wars.
Menggunakan data dari SWAPI untuk membuat game trivia Star Wars.
Untuk mengakses SWAPI, developer perlu mendaftar akun dan memperoleh API key. Selanjutnya data dapat diambil melalui endpoint-endpoint yang telah disediakan oleh SWAPI. Data yang bisa diambil antara lain karakter, planet, film, dan jenis-jenis kendaraan dalam Star Wars. Data tersebut tersedia dalam format JSON yang dapat dengan mudah diintegrasikan dengan berbagai bahasa pemrograman seperti Python dan JavaScript.

Third Party Golang
Why Need Third Party for Network?
What and why Echo?
Install Echo
Basic Routes & Controller
Render Data (JSON Response)
Retrieve Data
Binding Data
Overview Final Project

THIRD PARTY LIBRARY GOLANG
Echo
go kit
logrus
gorm.io
cobra

Binding data merupakan proses menghubungkan data dari suatu sumber dengan tampilan pada aplikasi. Hal ini dilakukan agar data yang ditampilkan di aplikasi selalu terupdate secara otomatis berdasarkan perubahan data yang terjadi pada sumbernya.

Contoh penggunaan binding data adalah pada aplikasi web yang menampilkan data dari database. Dengan menggunakan teknik binding data, ketika data pada database berubah, maka tampilan data pada aplikasi web juga akan berubah secara otomatis tanpa harus melakukan refresh halaman.

Pada pengembangan aplikasi, binding data dapat diterapkan pada berbagai platform, baik itu aplikasi desktop, web, maupun mobile. Beberapa framework atau library yang biasanya digunakan untuk mengimplementasikan teknik binding data seperti AngularJS, ReactJS, VueJS, EmberJS, dan lain sebagainya.
