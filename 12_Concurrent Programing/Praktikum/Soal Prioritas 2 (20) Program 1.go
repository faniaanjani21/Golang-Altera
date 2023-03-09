// Hitung frekuensi huruf dalam teks menggunakan perhitungan paralel (Bersamaan).

// contoh output
// input :
// bekasi merupakan ibukota provinsi jawa barat
// output :
// e : 4
// i : 4
// o : 2
// t : 2

package main

import (
	"fmt"
	"strings"
)

func countChar(s string) map[rune]int {
	// inisialisasi map untuk menyimpan jumlah kemunculan setiap karakter
	charCount := make(map[rune]int)

	// konversi string menjadi slice rune agar dapat diiterasi karakter per karakter
	for _, r := range strings.ToLower(s) {
		// abaikan spasi dan karakter selain huruf
		if r == ' ' || !('a' <= r && r <= 'z') {
			continue
		}
		charCount[r]++
	}

	return charCount
}

func main() {
	s := "Bekasi merupakan ibukota provinsi Jawa Barat"
	charCount := countChar(s)

	for char, count := range charCount {
		fmt.Printf("%c : %d\n", char, count)
	}
}
