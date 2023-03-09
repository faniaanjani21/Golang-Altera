// Buatlah program di Golang yang dapat mengurutkan barang berdasarkan jumlah kemunculannya. Jika ada barang yang duplicate kamu hanya perlu memunculkan sekali, namun kamu perlu menampilkan total kemunculan barang tersebut!

// Sample Test Cases
// Input: ["js". "js". "golang". "ruby", "ruby". "js". "js"]
// Output: golang->1 ruby->2 js->4

package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

// Fungsi MostAppearItem menerima input array string dan mengembalikan slice dari tipe pair
// yang berisi nama dan jumlah kemunculan setiap item di dalam array.
func MostAppearItem(items []string) []pair {
	itemCounts := make(map[string]int)
	for _, item := range items {
		if _, ok := itemCounts[item]; ok {
			itemCounts[item]++
		} else {
			itemCounts[item] = 1
		}
	}

	pairs := make([]pair, len(itemCounts))
	i := 0
	for k, v := range itemCounts {
		pairs[i] = pair{name: k, count: v}
		i++
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	return pairs
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// [{js 4} {ruby 2} {golang 1}]
	fmt.Println(MostAppearItem([]string{"B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// [{A 4} {B 3} {D 2} {C 1}]
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
	// [{football 1} {basketball 1} {tenis 1}]
}
