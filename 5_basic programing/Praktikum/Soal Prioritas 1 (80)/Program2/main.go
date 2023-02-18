// buatlah sebuah program untuk menentukan apakah sebuah bilang adalah bilang ganjil atau genap

package main

import "fmt"

func main() {
	var num int

	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Printf("%d adalah bilangan genap.\n", num)
	} else {
		fmt.Printf("%d adalah bilangan ganjil.\n", num)
	}
}

// ! Output
// PS C:\Users\R\Desktop\Praktikum Basic Programming\Soal Prioritas 1 (80)> go run "c:\Users\R\Desktop\Praktikum Basic Programming\Soal Prioritas 1 (80)\main.go"
// Masukkan sebuah bilangan: 10
// 10 adalah bilangan genap.
// PS C:\Users\R\Desktop\Praktikum Basic Programming\Soal Prioritas 1 (80)> go run "c:\Users\R\Desktop\Praktikum Basic Programming\Soal Prioritas 1 (80)\main.go"
// Masukkan sebuah bilangan: 9
// 9 adalah bilangan ganjil.
