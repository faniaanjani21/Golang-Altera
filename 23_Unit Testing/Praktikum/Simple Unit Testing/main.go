package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// fungsi untuk melakukan operasi penjumlahan
func addition(vara, varb float64) float64 {
	return vara + varb
}

// fungsi untuk melakukan operasi pengurangan
func subtraction(vara, varb float64) float64 {
	return vara - varb
}

// fungsi untuk melakukan operasi perkalian
func multiplication(vara, varb float64) float64 {
	return vara * varb
}

// fungsi untuk melakukan operasi pembagian
func division(vara, varb float64) float64 {
	return vara / varb
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// meminta input dari user
	fmt.Print("Masukkan nilai nomor 1: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	nilainomor1, _ := strconv.ParseFloat(input1, 64)

	fmt.Print("Masukkan nilai nomor 2: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	nilainomor2, _ := strconv.ParseFloat(input2, 64)

	// contoh penggunaan fungsi operasi matematika
	fmt.Println("Hasil Penjumlahan: ", addition(nilainomor1, nilainomor2))
	fmt.Println("Hasil Pengurangan: ", subtraction(nilainomor1, nilainomor2))
	fmt.Println("Hasil Perkalian: ", multiplication(nilainomor1, nilainomor2))
	fmt.Println("Hasil Pembagian: ", division(nilainomor1, nilainomor2))
}
