// Buatlah program untuk menentukan substring yang sama di antara keduanya, dengan diberikan dua string A dan B!
package main

import (
	"fmt"
	"strings"
)

func compare(s1, s2 string) string {
	var result string
	for i := 0; i < len(s1); i++ {
		for j := i + 1; j <= len(s1); j++ {
			if strings.Contains(s2, s1[i:j]) {
				if len(s1[i:j]) > len(result) {
					result = s1[i:j]
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(compare("aka", "akashi"))
	fmt.Println(compare("kanggoro", "kang"))
	fmt.Println(compare("ki", "kijang"))
	fmt.Println(compare("kupu-kupu", "kupu"))
	fmt.Println(compare("ilalang", "ila"))
}

// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\9_String, Advance Function, Pointer, Struct, Method, Interface, Package & Error Handling> go run "c:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\9_String, Advance Function, Pointer, Struct, Method, Interface, Package & Error Handling\Praktikum\soal prioritas 1 (80)\program3\main.go"
// aka
// kang
// ki
// kupu
// ila
// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\9_String, Advance Function, Pointer, Struct, Method, Interface, Package & Error Handling>
