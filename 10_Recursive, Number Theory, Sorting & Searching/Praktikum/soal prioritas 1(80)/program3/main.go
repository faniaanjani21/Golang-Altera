// Dalam matematika, bilangan prima adalah bilangan asli yang lebih besar dari angka 1, yang faktor pembaginya adalah 1 dan bilangan itu sendiri. Angka 2 dan 3 adalah bilangan prima. Angka 4 bukan bilangan prima karena 4 bisa dibagi 2. Sepuluh deret bilangan prima yang pertama adalah 2, 3, 5, 7, 11, 13, 17, 19, 23 dan 29. Buatlah sebuah fungsi bernama getPrime yang menampilkan bilangan prima sesuai dengan deret urutannya.
// Sample Test Cases
//  Input: 1
// Output: 2
//  Input: 5
// Output: 11

package main

import (
	"fmt"
)

func primeX(number int) int {
	if number == 1 {
		return 2
	}

	count := 1
	candidate := 3
	for {
		isPrime := true
		for i := 2; i*i <= candidate; i++ {
			if candidate%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			count++
			if count == number {
				return candidate
			}
		}
		candidate += 2
	}
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
