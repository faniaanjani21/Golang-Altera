package main

import "fmt"

func main() {
	// slice untuk menyimpan data dinamis
	data := []int{1, 2, 3, 4, 5}
	fmt.Println("DATA PERTAMA:", data)

	// Mengubah data tanpa membuat slice baru
	for i := range data {
		data[i] *= 2
	}

	fmt.Println("Data yang telah di ubah", data)
}

// e complexity\Praktikum\program2\main.go"
// DATA PERTAMA: [1 2 3 4 5]
// Data yang telah di ubah [2 4 6 8 10]
// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\7_time complexity & space complexity>
