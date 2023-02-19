materi prblem solving
brute foce
divide conquer
gredy
dynaic programming

brute
mencoba semua kemungkinan

divide conquer
membagi masalah menjadi sub masalah

greedy
mencari solusi lokal

dynamic programming
mencari solusi global

materi ini berfungsi untuk mempercepat dalam pencarian dan memperkecil ruang pencarian

contoh program brute force
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

contoh program divide conquer
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
return n \* factorial(n-1)
}

contoh program greedy
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
for \_, item := range items {
if weight+item.Weight <= capacity {
value += item.Value
weight += item.Weight
} else {
remaining := capacity - weight
value += remaining \* item.Value / item.Weight
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

contoh program dynamic programming
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
