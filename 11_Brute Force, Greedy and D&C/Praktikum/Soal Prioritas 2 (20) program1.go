// Ada katak yang awalnya berada di atas Batu 1. Dia akan mengulangi tindakan berikut beberapa kali untuk mencapai batu N. Jika katak sedang berada di Batu i, lompat ke Batu i + 1 atau Batu i + 2. Di sini, biaya | hi - hj | terjadi, di mana j adalah batu untuk mendarat. Temukan biaya total minimum yang mungkin dikeluarkan sebelum katak mencapai Batu N.

package main

import (
	"fmt"
	"math"
)

func frog(jums []int, n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return int(math.Abs(float64(jums[0] - jums[1])))
	}
	return int(math.Min(float64(frog(jums, n-1)+int(math.Abs(float64(jums[n-1]-jums[n-2])))), float64(frog(jums, n-2)+int(math.Abs(float64(jums[n-1]-jums[n-3]))))))
}

func main() {
	fmt.Println(frog([]int{10, 30, 40, 20}, 4))         // 30
	fmt.Println(frog([]int{30, 10, 60, 10, 60, 50}, 6)) // 40
}
