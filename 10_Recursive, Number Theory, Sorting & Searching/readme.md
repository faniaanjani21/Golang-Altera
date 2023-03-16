# Pengenalan Algoritma

Algoritma merupakan langkah-langkah atau instruksi yang digunakan untuk menyelesaikan suatu masalah atau tugas. Algoritma dapat digunakan dalam berbagai bidang, seperti pemrograman, matematika, fisika, dan sebagainya.

## Recursive

Recursive merupakan fungsi yang memanggil dirinya sendiri secara berulang sampai mencapai kondisi dasar atau base case. Recursive sering digunakan dalam metode pencarian dan pengurutan data. Namun, penggunaan rekursi yang berlebihan dapat menyebabkan stack overflow dan performa yang buruk.

Contoh:

```go
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
```

## Number Theory

Number Theory merupakan cabang matematika yang mempelajari sifat-sifat bilangan bulat.

Contoh:

```go
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
```

## Searching

Searching merupakan metode untuk mencari data dalam suatu himpunan. Beberapa metode yang sering digunakan dalam searching antara lain linear search, binary search, dan hash table.

Contoh:

```go
package main

import "fmt"

func linearSearch(slice []int, target int) bool {
    for _, num := range slice {
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
```

## Sorting

Sorting merupakan metode untuk mengurutkan data dalam suatu himpunan. Beberapa metode yang sering digunakan dalam sorting antara lain bubble sort, selection sort, insertion sort, merge sort, dan quick sort.

Contoh:

```go
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
```
