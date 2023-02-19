// 1. Kamu diminta untuk membeli sebuah barang, dan tantangan kali ini kamu harus bisa membeli barang dengan jumlah maksimum dengan uan yang kamu miliki.

//     Program ini menerima `money` sebagai parameter pertama dan yang kedua berupa list harga produk yang dapat dibeli. Kamus harus menampilkan nilai jumlah barang yang bisa dibeli.

//     Sample Test Case

//     Input: `money = 50000`, `productPrice = [25000, 25000, 10000, 14000]`

//     Output: 3

//     Input: `money = 50000`, `productPrice = 15000, 10000, 12000, 5000, 3000]`

//     Output: 5

package main

import "fmt"

func MaximumBuyProduct(money int, productPrice []int) {
	count := 0
	for _, price := range productPrice {
		if price <= money {
			count++
			money -= price
		}
	}
	fmt.Println(count)
}

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})      // 3
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}) // 4
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})   // 4
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})           // 1
	MaximumBuyProduct(0, []int{10000, 30000})                        // 0
}

// 2
// 3
// 4
// 1
// 0
// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\10_Recursive, Number Theory, Sorting & Searching>
