// Buatlah program playingDomino yang menerima 2 parameter array; parameter pertama merupakan kartu domino yang ada di tangan, • Parameter kedua merupakan kartu yang sedang ada di deck. Jika ada kartu yang disarankan maka output: [x,y], jika tidak ada kartu yang sesuai maka keluarkan: 'tutup kartu'.

package main

import (
	"fmt"
)

func playingDomino(cards [][]int, deck []int) interface{} {
	for _, card := range cards {
		if card[0] == deck[0] || card[0] == deck[1] {
			return card
		}
		if card[1] == deck[0] || card[1] == deck[1] {
			return card
		}
	}
	return "tutup kartu"
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{3, 3}, []int{4, 3}}, []int{3, 4}))
	// [3 4]

	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
	// [3 6]

	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}, []int{5, 1}}, []int{4, 6}))
	// tutup kartu
}

// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\10_Recursireve, Number Theory,go run "c:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\10_Recursive, Number Theory, Sorting & Searching\Praktikum\Soal Prioritas 2 (20)\program1\main.go"(20)\pro
// [3 4]
// [6 5]
// [6 6]
// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\10_Recursive, Number Theory, Sorting & Searching>
