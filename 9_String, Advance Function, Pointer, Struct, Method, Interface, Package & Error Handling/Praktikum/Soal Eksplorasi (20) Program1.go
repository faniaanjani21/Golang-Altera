// Implementasikan interface yang terdiri dari method encode dan decode. Algoritma enkripsi yang dapat digunakan adalah subtitusi cipher.

// Sample Test Case

// Input:

// [1] Encrypt

// [2] Decrypt

// Choose your menu? 1

// Input Student Name: rizky

// Output

// Encode of student’s name rizky is irapb

// package main

// import "fmt"

// type student struct {
// 	name string
// 	nameEncode string
// 	score int
// }

// type Chiper interface {
// 	Encode() string
// 	Decode() string
// }

// func (s *student) Encode() string {
// 	var nameEncode = ""

// 	// your code here

// 	return nameEncode
// }

// func (s *student) Decode() string {
// 	var nameDecode = ""

//   // your code here

//   return nameDecode
// }

// func main() {
// 	var menu int
//   var a student = student{}
// 	var c Chiper = &a

// 	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
// 	fmt.Scan(&menu)

// 	if menu == 1 {
// 		fmt.Print("\nInput Student Name: ")
// 		fmt.Scan(&a.name)
// 		fmt.Print("\nEncode of student’s name " + a.name + "is : " + c.Encode())
// 	} else if menu == 2 {
// 		fmt.Print("\nInput Student Name: ")
// 		fmt.Scan(&a.name)
// 		fmt.Print("\nDecode of student’s name " + a.name + "is : " + c.Decode())
// 	}
// }

package main

import (
	"fmt"
	"strings"
)

type Chiper interface {
	Encode() string
	Decode() string
}

type student struct {
	name       string
	nameEncode string
}

func (s *student) Encode() string {
	var nameEncode strings.Builder

	for _, char := range s.name {
		if char >= 'a' && char <= 'z' {
			nameEncode.WriteByte('a' + byte(char-'a'+1)%26)
		} else {
			nameEncode.WriteByte(byte(char))
		}
	}

	s.nameEncode = nameEncode.String()
	return s.nameEncode
}

func (s *student) Decode() string {
	var nameDecode strings.Builder

	for _, char := range s.nameEncode {
		if char >= 'b' && char <= 'z' {
			nameDecode.WriteByte('a' + byte(char-'b'+25)%26)
		} else {
			nameDecode.WriteByte(byte(char))
		}
	}

	return nameDecode.String()
}

func main() {
	var menu int
	var a student

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + a.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of student’s name " + a.nameEncode + " is : " + a.Decode())
	}
}

// face, Package & Error Handling>  go run "c:\Users\R\Docum_String, Advance Function, Pointer, Struct, Method, Inteents\GitHub\go-muhammad-reza-hidayat\9_String, Advance Function, Pointer, Struct, Method, Interface, Package & Error Handling\Praktikum\Soal Eksplorasi (20)\main.go"
// [1] Encrypt
// [2] Decrypt
// Choose your menu? 1

// Input Student Name: reza

// Encode of student’s name reza is : sfab
// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\9_String, Advance Function, Pointer, Struct, Method, Interface, Package & Error Handling>
