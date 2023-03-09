// **Caesar Cipher**

// Buatlah sebuah method dengan parameter `offset` bertipe integer dan `input` bertipe string. Method ini menghasilkan sebuah string baru yang dimana setiap huruf dilakukan pergeseran berdasarkan `offset` yang merupakan jumlah pergeseran hurufnya. String `input` diasumsikan berisi huruf kecil saja dan spasi. Sebagai contoh, ketika terdapat huruf `z` yang digeser dengan `offset` sebesar 3 maka huruf hasil pergeseran adalah `c.`

// Daftar karakter ASCII dapat dilihat di link [berikut](https://id.wikipedia.org/wiki/ASCII).

// Berdasarkan referensi ASCII, huruf `a` memiliki kode 97, huruf `b` memiliki kode 98, `z` memiliki kode 122. Anda bisa menggunakan operator modulo jika diperlukan.

// Test Case 1

// Input: `offset = 3` , `input = abc`

// Output: `def`

// Test Case 2

// Input: `offset = 1`, `input = abcdefghijklmnopqrstuvwxyz`

// Output: `bcdefghijklmnopqrstuvwxyza`

// Test Case 3

// Input: `offset = 1000`, `input = abcdefghijklmnopqrstuvwxyz`

// Output: `mnopqrstuvwxyzabcdefghijkl`

// package main

// import "fmt"

// func caesar(offset int, input string) string {
// 	// your code here
// }

// func main() {
// 	fmt.Println(caesar(3, "abc")) // def
// 	fmt.Println(caesar(2, "alta")) // def
// 	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi
// 	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza
//   fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
// }

package main

import "fmt"

func caesar(offset int, input string) string {
	// inisialisasi variabel untuk menyimpan hasil enkripsi
	var encrypted string

	// lakukan perulangan untuk setiap karakter pada string input
	for _, char := range input {
		// jika karakter bukan spasi, lakukan enkripsi
		if char != ' ' {
			// konversi karakter menjadi kode ASCII
			charCode := int(char)
			// hitung kode ASCII setelah pergeseran
			shiftedCode := (charCode-97+offset)%26 + 97
			// konversi kode ASCII menjadi karakter
			shiftedChar := string(shiftedCode)
			// tambahkan karakter hasil enkripsi ke string encrypted
			encrypted += shiftedChar
		} else {
			// jika karakter adalah spasi, langsung tambahkan ke string encrypted
			encrypted += " "
		}
	}

	return encrypted
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmnowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}

// rface, Package & Error Handling> go run "c:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\9_String, Advance Function, Pointer, Struct, Method, Interface, Package & Error Handling\Praktikum\Soal Prioritas 2 (20)\main.go"
// def
// cnvc
// kvdobbkkmknowi
// bcdefghijklmnopqrstuvwxyza
// mnopqrstuvwxyzabcdefghijkl
// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\9_String, Advance Function, Pointer, Struct, Method, Interface, Package & Error Handling>
