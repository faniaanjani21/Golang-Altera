Recursive
Number Theory
Searching
Sorting

\*Recursive

funsgi recursive bisa di gunakan dalam metode pencarian dan metode pengurutan data

recursive merupakan perulangan atau bisa di sebut fungsi memanggil diri nya sendri seacara berulang ulang sampe ke dasar nya
contoh :

package main

import "fmt"

func fibonacci(n int) int {
if n <= 1 {
return n
}
return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
fmt.Println("10 bilangan pertama dalam deret Fibonacci:")
for i := 0; i < 10; i++ {
fmt.Print(fibonacci(i), " ")
}
}

penggunaan rekursi yang berlebihan dapat menyebabkan stack overflow (tumpukan panggilan fungsi terlalu dalam) dan performa yang buruk, sehingga sebaiknya digunakan dengan hati-hati dan hanya pada kasus yang memang membutuhkan rekursi.

fungsi memanggil dirinya sendiri untuk memecahkan masalah lebih kecil, dan kemudian menggabungkan hasilnya untuk memecahkan masalah asli.

\*Number Theory
Number Theory merupakan cabang matematika yang mempelajari sifat-sifat bilangan bulat

contoh :
package main

import (
"fmt"
"math/big"
)

func main() {
n := big.NewInt(123456789) // bilangan yang akan difaktorkan
fmt.Printf("Faktorisasi prima dari %d adalah:\n", n)
factors := big.NewInt(2) // inisialisasi faktor awal dengan 2

    // ulangi pembagian hingga tidak dapat dibagi lagi
    for n.Cmp(big.NewInt(1)) > 0 {
    	for n.Mod(n, factors).Cmp(big.NewInt(0)) == 0 {
    		n.Div(n, factors)
    		fmt.Print(factors, " ")
    	}
    	factors.Add(factors, big.NewInt(1))
    }

}

\*Searching
Searching merupakan metode pencarian data

contoh :
contoh program sederhana untuk pencarian angka di dalam slice menggunakan linear search pada bahasa Golang:
package main

import "fmt"

func linearSearch(slice []int, target int) bool {
for \_, num := range slice {
if num == target {
return true
}
}
return false
}

func main() {
// contoh slice
nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // pencarian angka 5
    if linearSearch(nums, 5) {
        fmt.Println("Angka 5 ditemukan di dalam slice")
    } else {
        fmt.Println("Angka 5 tidak ditemukan di dalam slice")
    }

    // pencarian angka 11
    if linearSearch(nums, 11) {
        fmt.Println("Angka 11 ditemukan di dalam slice")
    } else {
        fmt.Println("Angka 11 tidak ditemukan di dalam slice")
    }

}

\*Sorting
Sorting merupakan metode pengurutan data

contoh :
package main

import "fmt"

func bubbleSort(slice []int) {
n := len(slice)
for i := 0; i < n; i++ {
for j := 0; j < n-i-1; j++ {
if slice[j] > slice[j+1] {
slice[j], slice[j+1] = slice[j+1], slice[j]
}
}
}
}

func main() {
// contoh slice yang akan diurutkan
nums := []int{4, 2, 7, 1, 8, 3, 5, 6}

    // urutkan slice menggunakan bubble sort
    bubbleSort(nums)

    // cetak slice yang telah diurutkan
    fmt.Println(nums)

}
