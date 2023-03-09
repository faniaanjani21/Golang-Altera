// Buatlah sebuah method untuk menghitung perkiraan jarak yang bisa ditempuh berdasarkan bahan bakar yang sedang terisi (1.5 L / mill), method tersebut receiver kepada struct Car yang memiliki property type dan fuelIn

package main

import "fmt"

// defisnin struktur car
type Car struct {
	Type   string
	FuelIn float64
}

// defisnisi method untuk emnghitung jarak yang bisa ditempuh
func (car Car) CalculateRange() float64 {
	// 1.5 liter berpa mill
	return car.FuelIn / 0.0015
}

func main() {
	// membuat objek model sedan
	car := Car{Type: "Sedan", FuelIn: 45.0}

	// menghitung jarak yang bisa ditempuh
	rangeInKM := car.CalculateRange()

	// output nya
	fmt.Printf("The %s can travel up to %.2f kilometers with the current fuel\n", car.Type, rangeInKM)
}

// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\9_String, Advance Function, Pointer, Struct, Method, Interface, Package & Error Handling> go run "c:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\9_String, Advance Function, Pointer, Struct, Method, Interface, Package & Error Handling\Praktikum\soal prioritas 1 (80)\program1\main.go"
// The Sedan can travel up to 30000.00 kilometers with the current fuel
// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\9_String, Advance Function, Pointer, Struct, Method, Interface, Package & Error Handling>
