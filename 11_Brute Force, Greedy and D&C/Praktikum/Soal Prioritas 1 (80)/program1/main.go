// Diberi bilangan bulat n, kembalikan array ans dengan panjang n + 1 sehingga untuk setiap i (0 <= i <= n), ans[i] adalah bilangan 1 dalam representasi biner dari i
// Input: n = 2
// Output: [0,1,1]
// Explanation:
//  0 --> 0
//  1 --> 1
//  2 --> 10

package main

import "fmt"

func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = ans[i>>1] + i&1
	}
	return ans
}

func main() {
	fmt.Println(countBits(2)) // [0 1 1]
	fmt.Println(countBits(5)) // [0 1 1 2 1 2]
}
