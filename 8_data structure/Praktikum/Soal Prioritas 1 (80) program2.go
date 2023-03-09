package main

import "fmt"

func Mapping(slice []string) map[string]int {
	// Buat sebuah map untuk menyimpan setiap string yang sudah ditemukan
	counts := make(map[string]int)

	// Iterasi setiap elemen pada slice
	for _, s := range slice {
		// Tambahkan 1 ke nilai yang sudah ada pada map jika string sudah pernah muncul
		// Jika tidak, set nilai map ke 1
		counts[s] = counts[s] + 1
	}

	return counts
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	// map[adi:1 asd:2 qwe:3]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	// map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{}))
	// map[]
}

// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\8_data structure> go run "c:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\8_data structure\Praktikum\program2\main.go"
// map[adi:1 asd:2 qwe:3]
// map[asd:2 qwe:1]
// map[]
// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\8_data structure>
