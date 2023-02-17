// ? Buatlah sebuah program untuk memeriksa apakah sebuah kata adalah palindrome atau bukan, serta coba terapkan scanner untuk menangkap inputan dari console,

// contoh output
// apakah palindrome
// masukan kata : aya aya
// captured : aya aya
// palindrome

package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Masukkan kata: ")
	fmt.Scanln(&input)

	// Menghapus spasi dan mengubah ke lowercase
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ToLower(input)

	// Memeriksa apakah kata adalah palindrome
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			fmt.Println("Bukan palindrome")
			return
		}
	}

	fmt.Println("Palindrome")
}

// ! output
// PS C:\Users\R> go run "c:\Users\R\Desktop\Praktikum Basic Programming\Soal Prioritas 2 (20)\mian.go"
// Masukkan angka: 26
// Faktor bilangan dari26adalah: 1 2 13 26
