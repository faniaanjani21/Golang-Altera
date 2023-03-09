// Buatlah program di Golang untuk menemukan nilai maksimum serta minimum di antara 6 angka inputan. Gunakan multiple return fungsi, pointer untuk referencing maupun deferencing!

// sample test case
// input :
// 1
// 2
// 3
// 9
// 7
// 8
// output
// 9 is max
// 1 is min

package main

import "fmt"

func getMinMax(numbers ...*int) (min int, max int) {
	// inisialisasi min dan max dengan nilai pertama pada array
	min = *numbers[0]
	max = *numbers[0]

	// loop melalui array, membandingkan nilai setiap elemen dengan min dan max
	for _, num := range numbers {
		if *num < min {
			min = *num
		}
		if *num > max {
			max = *num
		}
	}
	return
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	// pass alamat variabel ke fungsi getMinMax menggunakan operator &
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Printf("%d is min\n", min)
	fmt.Printf("%d is max\n", max)
}

// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\9_String, Advance Function, Pointer, Struct, Method, Interface, Package & Error Handling> go run "c:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\9_String, Advance Function, Pointer, Struct, Method, Interface, Package & Error Handling\Praktikum\soal prioritas 1 (80)\program4\main.go"
// 1
// 2
// 3
// 10
// 2
// 5
// 1 is min
// 10 is max
// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\9_String, Advance Function, Pointer, Struct, Method, Interface, Package & Error Handling>
