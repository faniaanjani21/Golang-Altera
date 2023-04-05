cara menjalankan nya
go build init login
Perintah go build dan go init adalah dua perintah terpisah dalam Go. Kedua perintah tersebut memiliki fungsi yang berbeda.

go build adalah perintah yang digunakan untuk mengumpulkan package Go dan menjadikannya sebagai binary file (eksekusi program). Dalam hal ini, ketika Anda menulis go build, Go akan mengumpulkan semua file .go yang ada di direktori saat ini dan mengeksekusinya menjadi sebuah binary file yang bisa Anda jalankan. Namun, pada kasus penggunaan Anda (go build init login), itu tidak valid karena init dan login bukanlah argumen yang valid atau spesifik bagi perintah go build.

Sedangkan go init adalah perintah yang digunakan untuk membuat sebuah module Go baru atau untuk menambahkan file go ke dalam module yang sudah dibuat sebelumnya. Modul Go adalah struktur direktori yang berisi file-file Go dan memiliki file go.mod yang bertanggung jawab untuk mendefinisikan modul tersebut serta dependensinya. Namun, menggunakan go init sendirian juga tidak valid karena Anda harus menentukan nama module yang ingin Anda buat. Sebagai contoh: go mod init nama-module-anda.

Jadi, jika Anda ingin membuat sebuah module dengan nama login dan menambah beberapa file ke dalamnya, maka langkah-langkah yang dapat dilakukan adalah:

Ketikkan perintah go mod init login untuk membuat sebuah modul baru bernama login.
Setelah modul berhasil dibuat, buatlah beberapa file .go yang relevan dan tambahkan konfigurasi package yang diperlukan ke dalam file-file tersebut.
Ketika Anda ingin mengumpulkan semua file .go dalam modul login, kemudian menjadikannya binary file, gunakanlah perintah go build.
Semoga penjelasan ini dapat membantu!

go mod tidy
Perintah go mod tidy adalah salah satu perintah pada Go yang berguna untuk menghapus module yang tidak terpakai dari go.mod dan menambahkan module baru ke dalam go.mod jika diperlukan.

Perintah ini akan mengecek dependensi apa saja yang digunakan oleh package Anda di project Go. Jika ada dependensi yang tidak lagi digunakan, dependensi tersebut akan dihapus dari file go.mod dan go.sum. Perintah ini juga akan memeriksa dependencies yang diperlukan saat menjalankan aplikasi, kemudian menambahkan atau mengupdate dependensi yang hilang atau belum ada.

Dalam pengembangan proyek Go, biasanya dependensi sering berubah-ubah tergantung dari kebutuhan. Sehingga, dengan menggunakan perintah go mod tidy, Anda dapat dengan mudah menyeimbangkan dependensi dalam modul Anda yang dapat membantu Anda mencegah konflik dependency ketika melakukan build atau kompilasi.

Jadi, secara umum, perintah go mod tidy sangat membantu dalam mendapatkan file dependency yang efisien dan mengelola daftar depedency yang lebih mudah.
