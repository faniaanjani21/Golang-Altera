// Angka Fibonacci adalah serangkaian angka di mana setiap angka adalah jumlah dari keduanya nomor sebelumnya. Beberapa angka Fibonacci pertama adalah: 0, 1, 1, 2, 3, 5, 8, ….
// Buatlah fungsi untuk menghitung angka Fibonacci ke-n (top-down)!

// spmle test casesinput : 5
// ouput : 5

package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(fib(0))  // 0
	fmt.Println(fib(1))  // 1
	fmt.Println(fib(2))  // 1
	fmt.Println(fib(3))  // 2
	fmt.Println(fib(5))  // 5
	fmt.Println(fib(6))  // 8
	fmt.Println(fib(7))  // 13
	fmt.Println(fib(9))  // 34
	fmt.Println(fib(10)) // 55
}
