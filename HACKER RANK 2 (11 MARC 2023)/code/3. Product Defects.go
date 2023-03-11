package main

import (
	"fmt"
)

func findLargestSquareSize(rezasamples [][]int) int {
	n := len(rezasamples)
	maxSize := 0
	// membuat tabel dp untuk menyimpan ukuran sub-matriks persegi terbesar
	// produk cacat yang berakhir pada setiap posisi (i,j) dalam matriks sampel
	dp := make([][]int, n)
	for p := range dp {
		dp[p] = make([]int, n)
	}

	// menginisialisasi baris pertama dan kolom pertama dari tabel dp
	for p := 0; p < n; p++ {
		dp[p][0] = rezasamples[p][0]
		dp[0][p] = rezasamples[0][p]
	}

	// menghitung ukuran sub-matriks persegi terbesar dari produk cacat
	// berakhir pada setiap posisi (i,j) dalam matriks sampel
	for p := 1; p < n; p++ {
		for b := 1; b < n; b++ {
			if rezasamples[p][b] == 1 {
				dp[p][b] = 1 + min(dp[p-1][b-1], min(dp[p-1][b], dp[p][b-1]))
				if dp[p][b] > maxSize {
					maxSize = dp[p][b]
				}
			}
		}
	}

	return maxSize
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// mendapatkan masukan
	var n int
	fmt.Scan(&n)
	rezasamples := make([][]int, n)
	for p := range rezasamples {
		rezasamples[p] = make([]int, n)
		for b := range rezasamples[p] {
			fmt.Scan(&rezasamples[p][b])
		}
	}

	//  menemukan ukuran sub-matriks persegi terbesar dari produk cacat
	maxSize := findLargestSquareSize(rezasamples)

	// mencetak hasilnya
	fmt.Println(maxSize)
}
