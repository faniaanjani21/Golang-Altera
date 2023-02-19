// 1. Terdapat dua bilangan integer yaitu x dan n. Buatlah sebuah fungsi untuk melakukan perhitungan perpangkatan (x^n dibaca x pangkat n). Time complexity dari sebuah fungsi perpangkatan harus lebih cepat dari O(n). Contoh time complexity yang optimal adalah logaritmik.

//     **Sample Test Cases**

//     Input : x = 2, n = 3

//     Output : 8

//     Input : x = 7, n = 2

//     Output : 49

package main

import (
	"fmt"
)

func power(x, n int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		p := power(x, n/2)
		return p * p
	}
	return x * power(x, n-1)
}

func main() {
	x := 2
	n := 3
	result := power(x, n)
	fmt.Printf("%d pangkat %d = %d\n", x, n, result)

	x = 7
	n = 2
	result = power(x, n)
	fmt.Printf("%d pangkat %d = %d\n", x, n, result)
}

// complexity & space complexity\Praktikum\program4-tugas\main.go"
// 2 pangkat 3 = 8
// 7 pangkat 2 = 49
