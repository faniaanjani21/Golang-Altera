// Buatlah sebuah fungsi yang mencetak angka kelipatan x, dimana x merupakan parameter bilangan bulat positif. lalu jalankan fungsi tersebut dengan menerapkan goroutine, dengan interval waktu 3 detik!

package main

import (
	"fmt"
	"time"
)

func printMultiples(x int) {
	for i := 1; ; i++ {
		if i%x == 0 {
			fmt.Println(i)
		}
		time.Sleep(3 * time.Second)
	}
}

func main() {
	go printMultiples(3)
	time.Sleep(30 * time.Second)
}
