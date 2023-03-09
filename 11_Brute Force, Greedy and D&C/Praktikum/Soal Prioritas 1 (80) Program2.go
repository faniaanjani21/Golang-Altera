// Diberi bilangan bulat numRows, kembalikan numRows pertama dari segitiga Pascal. Dalam segitiga Pascal, setiap angka adalah jumlah dari dua angka tepat di atasnya seperti yang ditunjukkan:

// input: numRows = 5
// output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

package main

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		ans[i] = make([]int, i+1)
		ans[i][0], ans[i][i] = 1, 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}
	return ans
}
