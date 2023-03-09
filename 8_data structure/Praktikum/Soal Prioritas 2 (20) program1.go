package main

import "fmt"

func Pairsum(arr []int, target int) []int {
	i := 0
	j := len(arr) - 1
	for i < j {
		if arr[i]+arr[j] == target {
			return []int{i, j}
		} else if arr[i]+arr[j] < target {
			i++
		} else {
			j--
		}
	}
	return []int{}
}

func main() {
	fmt.Println(Pairsum([]int{1, 2, 3, 4, 5}, 6)) // [1 3]
	fmt.Println(Pairsum([]int{2, 5, 9, 11}, 11))  // [0 2]
	fmt.Println(Pairsum([]int{1, 3, 5, 7}, 12))   // [1 3]
	fmt.Println(Pairsum([]int{1, 4, 6, 8}, 10))   // [1 2]
	fmt.Println(Pairsum([]int{1, 5, 6, 7}, 6))    // [0 1]
}

// Diberi array angka yang diurutkan dan jumlah target, temukan pasangan dalam array yang jumlahnya sama dengan target yang diberikan. Tulis fungsi untuk mengembalikan indeks dari dua angka (yaitu pasangan) sehingga jika value pada index tersebut dijumlahkan sesuai target yang diberikan.

// sample test
// input : [12345], target = 6
// explanation the numbers at index 1 and 3 add up to 6: 2 + 4 = 6, so the answer is [1,3]

// input : [25911], target = 11
// output : [0,2]
// explanation: the numbers at index 0 and 2 add up to 11: 2 + 9 = 11, so the answer is [0,2]

// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\8_data structure> go run "c:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\8_data structure\Praktikum\Soal Prioritas 2 (20)\program1\main.go"
// [0 4]
// [0 2]
// [2 3]
// [1 2]
// [0 1]
// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\8_data structure>
