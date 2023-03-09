// ? Buatlah sebuah program untuk mencetak segitiga asterik seperti dibawah ini!

package main

import "fmt"

func main() {
	var tinggi int

	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scan(&tinggi)

	for i := 1; i <= tinggi; i++ {
		for j := 1; j <= tinggi-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// ! output
// PS C:\Users\R> go run "c:\Users\R\Desktop\Praktikum Basic Programming\Soal Prioritas 2 (20)\mian.go"
// Masukkan tinggi segitiga: 5
//     *
//    ***
//   *****
//  *******
// *********
