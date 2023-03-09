package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	// Buat sebuah map untuk menyimpan jumlah kemunculan setiap angka pada string
	counts := make(map[rune]int)

	// Iterasi setiap karakter pada string dan hitung jumlah kemunculannya
	for _, r := range angka {
		counts[r]++
	}

	// Buat slice untuk menyimpan angka yang hanya muncul sekali
	var result []int

	// Iterasi setiap karakter pada string lagi dan tambahkan ke slice jika kemunculannya adalah 1
	for _, r := range angka {
		if counts[r] == 1 {
			n, err := strconv.Atoi(string(r))
			if err == nil {
				result = append(result, n)
			}
		}
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}

// sample test case
// input : "76523752"
// output : [6 3]

// input : "1122"
// output : []

// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\8_data structure> go run "c:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\8_data structure\Praktikum\program3\main.go"
// [4]
// [6 3]
// [1 2 3 4 5]
// []
// [8 7 2 5 4]
// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\8_data structure>
