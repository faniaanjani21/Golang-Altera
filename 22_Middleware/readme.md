# Apa itu Middleware?

Middleware adalah perangkat lunak yang bertindak sebagai jembatan antara aplikasi dan sistem operasi atau infrastruktur yang mendukung aplikasi. Middleware membantu memfasilitasi komunikasi dan pertukaran data antara aplikasi yang berbeda, serta menyediakan fungsi tambahan seperti manajemen sesi, autentikasi, dan caching.

Contoh middleware meliputi server aplikasi, message broker, gateway API, web server, dan banyak lagi. Middleware sangat penting dalam lingkungan komputasi terdistribusi di mana beberapa aplikasi atau sistem harus bekerja bersama untuk mencapai tujuan yang sama. Dengan memungkinkan aplikasi berkomunikasi satu sama lain dan mengakses sumber daya yang sama, middleware memungkinkan pengembangan dan operasi sistem yang lebih efisien dan terkoordinasi.

# Contoh Program Golang Middleware

Berikut adalah contoh program Golang yang menggunakan middleware untuk melakukan logging setiap kali request HTTP diterima:

```go
package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	// Membuat handler untuk endpoint "/hello"
	helloHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	// Membuat middleware untuk logging
	loggingMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Printf("[%s] %s %s", r.Method, r.URL.Path, time.Since(start))
		})
	}

	// Menggabungkan handler dan middleware
	handler := loggingMiddleware(helloHandler)

	// Menjalankan server HTTP di port 8080
	log.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", handler)
	log.Fatal(err)
}
```

Dalam contoh ini, ketika server HTTP dijalankan, program akan mencetak log setiap kali request
