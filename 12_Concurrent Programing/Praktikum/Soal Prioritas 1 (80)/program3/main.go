// Buatlah program yang mencetak bilangan kelipatan 3 dengan menerapkan goroutine dan buffer channel!

package main

import "fmt"

func main() {
	// membuat channel dengan buffer 10
	ch := make(chan int, 10)

	// menjalankan goroutine untuk mengirimkan bilangan kelipatan 3 ke channel
	go func() {
		for reza := 1; ; reza++ {
			if reza%3 == 0 {
				ch <- reza
			}
		}
	}()

	// melakukan pembacaan dari channel
	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}
}
