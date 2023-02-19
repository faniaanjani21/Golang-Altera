// Buatlah  program yang yang menerapkan mutex! Jenis program yang dibuat bebas (contoh: perhitungan faktorial).

package main

import (
	"fmt"
	"sync"
)

var (
	za = 10
	mu sync.Mutex
)

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(za)

	for i := 1; i <= za; i++ {
		go func(num int) {
			defer wg.Done()
			mu.Lock()
			result := factorial(num)
			fmt.Printf("Factorial of %d is %d\n", num, result)
			mu.Unlock()
		}(i)
	}
	wg.Wait()
}
