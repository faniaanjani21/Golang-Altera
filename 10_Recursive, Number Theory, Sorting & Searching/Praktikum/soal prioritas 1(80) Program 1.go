// - Dalam matematika, bilangan Fibonacci adalah barisan yang didefinisikan secara rekursif sebagai berikut:

//     Penjelasan: barisan ini berawal dari 0 dan 1, kemudian angka berikutnya didapat dengan cara menambahkan kedua bilangan yang berurutan sebelumnya. Dengan aturan ini, maka barisan bilangan Fibonacci yang pertama adalah:

//     0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946…..

package main

import "fmt"

func fibonacci(number int) int {
	if number == 0 {
		return 0
	} else if number == 1 {
		return 1
	} else {
		return fibonacci(number-1) + fibonacci(number-2)
	}
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(19)) // 4181
	fmt.Println(fibonacci(12)) // 144
}

// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\10_Recursive, Number Theory, Sorting & Searching> go run
// "c:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\10_Recursive, Number Theory, Sorting & Searching\Praktikum\program1\main.go"
// 0
// 1
// 34
// 4181
// 144
// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\10_Recursive, Number Theory, Sorting & Searching>
