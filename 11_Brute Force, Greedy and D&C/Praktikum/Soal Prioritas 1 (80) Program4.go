// - Kamu memiliki tiga bilangan bulat yang berbeda, x, y dan z, yang memenuhi tiga hubungan berikut:
//     - x + y + z = A
//     - xyz = B
//     - x^2 + y^2 + z^2 = C

//     kamu diminta untuk menulis sebuah program yang memecahkan x, y dan z untuk nilai yang diberikan A, B dan C. (1 ≤ A, B, C ≤ 10000).

// Sample Test Cases
// Input: 123
// Output: No solution.
// Input: 6 6 14
// Output: 123

package main

import "fmt"

func simplequadratic(a, b, c int) (int, int, int) {
	x := (a + b + c) / 3
	y := (a * b * c) / x
	z := (a*a + b*b + c*c) / (3 * x * x)
	return x, y, z
}
func main() {
	fmt.Println(simplequadratic(123, 0, 0)) // No solution.
	fmt.Println(simplequadratic(6, 6, 14))  // 123
}
