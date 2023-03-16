# Materi Problem Solving

Problem Solving adalah suatu proses untuk menyelesaikan suatu masalah dengan cara yang sistematis dan logis. Terdapat beberapa teknik yang digunakan dalam problem solving, di antaranya adalah brute force, divide conquer, greedy, dan dynamic programming.

## Brute Force

Brute Force adalah teknik yang mencoba semua kemungkinan solusi untuk menemukan solusi yang paling optimal. Teknik ini cukup sederhana dan mudah dipahami, tetapi dapat memakan waktu yang lama dan tidak efisien untuk masalah yang kompleks.

Contoh program Brute Force:

```go
package main

import "fmt"

func findMax(a int, b int) int {
    var max int = a
    if b > a {
        max = b
    }
    return max
}

func main() {
    fmt.Println(findMax(10, 20)) // Output: 20
    fmt.Println(findMax(0, -5)) // Output: 0
    fmt.Println(findMax(100, 100)) // Output: 100
}
```

## Divide Conquer

Divide Conquer adalah teknik yang memecah masalah menjadi sub-masalah yang lebih kecil dan kemudian menyelesaikan sub-masalah tersebut secara terpisah. Setelah semua sub-masalah terselesaikan, solusi akhir dapat digabungkan dari solusi sub-masalah.

Contoh program Divide Conquer:

```go
package main

import "fmt"

func main() {
    n := 5
    result := factorial(n)
    fmt.Printf("Factorial of %d is %d\n", n, result)
}

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
```

## Greedy

Greedy adalah teknik yang mencari solusi yang paling optimal pada setiap tahap dengan harapan bahwa solusi lokal tersebut akan menghasilkan solusi yang optimal secara keseluruhan. Teknik ini sederhana dan cepat, tetapi tidak selalu menghasilkan solusi yang paling optimal.

Contoh program Greedy:

```go
package main

import (
    "fmt"
    "sort"
)

type Item struct {
    Value, Weight int
}

type ByDensity []Item

func (a ByDensity) Len() int { return len(a) }
func (a ByDensity) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByDensity) Less(i, j int) bool { return a[i].Value/a[i].Weight > a[j].Value/a[j].Weight }

func MaxValue(items []Item, capacity int) int {
    sort.Sort(ByDensity(items))
    var value, weight int
    for _, item := range items {
        if weight+item.Weight <= capacity {
            value += item.Value
            weight += item.Weight
        } else {
            remaining := capacity - weight
            value += remaining * item.Value / item.Weight
            break
        }
    }
    return value
}

func main() {
    items := []Item{
        {60, 10},
        {100, 20},
        {120, 30},
    }
    capacity := 50
    maxValue := MaxValue(items, capacity)
    fmt.Printf("Maximum value for capacity %d: %d\n", capacity, maxValue)
}
```

## Dynamic Programming

Dynamic Programming adalah teknik yang menggunakan memoization untuk menyimpan hasil dari sub-masalah yang sudah dipecah sebelumnya dan menghindari perhitungan ulang yang tidak perlu. Teknik ini sangat efisien untuk masalah yang kompleks.

Contoh program Dynamic Programming:

```go
package main

import "fmt"

func Fibonacci(n int) int {
    fib := make([]int, n+1)
    fib[0] = 0
    fib[1] = 1

    for i := 2; i <= n; i++ {
        fib[i] = fib[i-1] + fib[i-2]
    }

    return fib[n]

}

func main() {
    fmt.Println(Fibonacci(10)) // Output: 55
}
```

Materi ini berfungsi untuk mempercepat dalam pencarian dan memperkecil ruang pencarian.
