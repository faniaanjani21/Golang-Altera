// - Tulis program untuk mengkonversi dari angka normal ke Angka Romawi!

//     Input: 4

//     Output: IV

//     Input: 9

//     Output: IX

//     Input: 23

//     Output: XXIII

// Input: 2021

// Output: MMXXI

// Input: 1646

// Output: MDCXLVI

package main

import "fmt"

func intToRoman(num int) string {
	// Array of roman symbols
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	// Array of values
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	// Create an empty string to hold the result
	result := ""

	// Loop through the values
	for i := 0; i < len(values); i++ {
		// Check if the current value is less than or equal to the num
		for values[i] <= num {
			// Add the roman symbol to the result
			result += symbols[i]

			// Subtract the value from the num
			num -= values[i]
		}
	}

	// Return the result
	return result
}

func main() {
	// Test cases
	fmt.Println(intToRoman(4))    // IV
	fmt.Println(intToRoman(9))    // IX
	fmt.Println(intToRoman(23))   // XXIII
	fmt.Println(intToRoman(2021)) // MMXXI
	fmt.Println(intToRoman(1646)) // MDCXLVI
}
