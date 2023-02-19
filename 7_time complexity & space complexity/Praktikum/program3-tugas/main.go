// 1. Dalam matematika, bilangan prima adalah bilangan asli yang lebih besar dari angka 1, yang faktor pembaginya adalah 1 dan bilangan itu sendiri. 2 dan 3 adalah bilangan prima. 4 bukan bilangan prima karena 4 bisa dibagi 2. Kamu diminta untuk membuat fungsi untuk menentukan bahwa sebuah bilangan termasuk bilangan prima atau tidak.

//     Buatlah solusi yang lebih optimal, dengan kompleksitas lebih cepat dari O(n) / O(n/2).

//     ********************************Sample Test Case********************************

//     Input: 1000000007

//     Output: Bilangan Prima

//     Input: 1500450271

//     Output: Bilangan Prima

package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	sqrtN := int(math.Sqrt(float64(n)))

	for i := 5; i <= sqrtN; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func main() {
	n := 1000000007
	if isPrime(n) {
		fmt.Printf("%d adalah bilangan prima.\n", n)
	} else {
		fmt.Printf("%d bukan bilangan prima.\n", n)
	}
}

// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\7_time complexity & space complexity\Praktikum\program3-tugas> go run main.go
// 1000000007 adalah bilangan prima.
