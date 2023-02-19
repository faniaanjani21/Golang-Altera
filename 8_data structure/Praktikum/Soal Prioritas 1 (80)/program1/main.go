package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	// Buat sebuah map untuk menyimpan setiap nama yang sudah muncul
	names := make(map[string]bool)

	// Buat sebuah slice baru untuk menyimpan hasil gabungan
	merged := make([]string, 0)

	// Tambahkan semua elemen dari arrayA ke slice merged
	for _, name := range arrayA {
		if _, exists := names[name]; !exists {
			names[name] = true
			merged = append(merged, name)
		}
	}

	// Tambahkan semua elemen dari arrayB ke slice merged
	for _, name := range arrayB {
		if _, exists := names[name]; !exists {
			names[name] = true
			merged = append(merged, name)
		}
	}

	return merged
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]
	fmt.Println(ArrayMerge([]string{"devil jin", "sergei"}, []string{}))
	// ["devil jin", "sergei"]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}

// PS C:\Users\R> go run "c:\Users\R\Desktop\main.go"
// [king devil jin akuma eddie steve geese]
// [sergei jin steve bryan]
// [alisa yoshimitsu devil jin law]
// [devil jin sergei]
// [hwoarang]
// []
// PS C:\Users\R>
