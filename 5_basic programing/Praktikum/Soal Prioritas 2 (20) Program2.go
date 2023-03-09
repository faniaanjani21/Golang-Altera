// buatlah sebuah program untuk mencetak faktor bilangan dari sebuah angka,
// input 26
// output 1 , 2 ,13, 26

package main

import "fmt"

func main() {
	var angka int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	fmt.Print("Faktor bilangan dari", angka, "adalah: ")

	for i := 1; i <= angka; i++ {
		if angka%i == 0 {
			fmt.Print(i, " ")
		}
	}

	fmt.Println()
}

// ! output
// PS C:\Users\R> go run "c:\Users\R\Desktop\Praktikum Basic Programming\Soal Prioritas 2 (20)\mian.go"
// Masukkan angka: 26
// Faktor bilangan dari26adalah: 1 2 13 26
