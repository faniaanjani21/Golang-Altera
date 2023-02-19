// Buatlah program yang mencetak bilangan kelipatan 3 dengan menerapkan goroutine dan channel!

package main

import (
	"fmt"
)

func main() {
	// membuat channel untuk mengirimkan data bilangan kelipatan 3
	ch := make(chan int)

	// menjalankan goroutine untuk mengirimkan bilangan kelipatan 3 ke channel
	go func() {
		for i := 1; ; i++ {
			if i%3 == 0 {
				ch <- i
			}
		}
	}()

	// melakukan pembacaan dari channel
	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}
}
