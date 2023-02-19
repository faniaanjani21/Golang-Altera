package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// quicksort dengan pivot pertama
	pivot := arr[0]
	var left []int
	var right []int
	var equal []int
	for _, x := range arr {
		if x < pivot {
			left = append(left, x)
		} else if x > pivot {
			right = append(right, x)
		} else {
			equal = append(equal, x)
		}

		// quicksort dengan pivot terakhir
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, equal...), right...)
}

// main function untuk menjalankan program
func main() {
	arr := []int{5, 1, 3, 6, 8, 2, 7, 4}
	fmt.Println("Sebelum diurutkan:", arr)
	arr = quickSort(arr)
	fmt.Println("Setelah diurutkan:", arr)
}

// e complexity\Praktikum\program1\main.go"
// Sebelum diurutkan: [5 1 3 6 8 2 7 4]
// Setelah diurutkan: [1 2 3 4 5 6 7 8]
