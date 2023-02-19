// Diberi matriks persegi, buatlah program untuk menghitung selisih absolut antara jumlah diagonalnya
// Diagonal kiri ke kanan : 1+5+9 = 15 . Diagonal kanan ke kiri : 3+5+9 = 17 . Perbedaan mutlak adalah |15 - 17| = 2.
// 123
// 456
// 989

package main

import "fmt"

func main() {
	fmt.Println(DiagonalDifference([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	})) // 2
	fmt.Println(DiagonalDifference([][]int{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	})) // 15
}

func DiagonalDifference(arr [][]int) int {
	n := len(arr)
	var leftDiagonalSum, rightDiagonalSum int
	for i := 0; i < n; i++ {
		leftDiagonalSum += arr[i][i]
		rightDiagonalSum += arr[i][n-i-1]
	}
	return abs(leftDiagonalSum - rightDiagonalSum)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
