// buatlah sebuah program untuk menghitung luas trapesium

package main

import "fmt"

func main() {
	var alas, tinggi, sisiMiring float64

	fmt.Print("Masukkan panjang alas: ")
	fmt.Scan(&alas)

	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&tinggi)

	fmt.Print("Masukkan panjang sisi miring: ")
	fmt.Scan(&sisiMiring)

	luas := (alas + sisiMiring) * tinggi / 2

	fmt.Printf("Luas trapesium adalah: %.2f\n", luas)
}

// ! Output
// PS C:\Users\R\Desktop\Praktikum Basic Programming\Soal Prioritas 1 (80)> go run "c:\Users\R\Desktop\Praktikum Basic Programming\Soal Prioritas 1 (80)\main.go"
// Masukkan panjang alas: 12
// Masukkan tinggi: 10
// Masukkan panjang sisi miring: 5
// Luas trapesium adalah: 85.00
// PS C:\Users\R\Desktop\Praktikum Basic Programming\Soal Prioritas 1 (80)>
