// Buat sebuah struct dengan nama Student yang mempunyai properti name dan score dalam bentuk slice kemudian simpan data siswa sebanyak 5 siswa. Setelah 5 siswa dimasukkan maka program menunjukkan skor rata-rata, siswa yang memiliki skor minimum dan maksimal? (implementasikan method)

// sample
// input :
// input 1 student name rizky
// input 1 student score 80
// input 2 student name kobar
// input 2 student score 75
// input 3 student name ismail
// input 3 student score 70
// input 4 student name umam
// input 4 student score 75
// input 5 student name yopan
// input 5 student score 60

// output :
// average scroe :72
// min score of student :yopan 60
// max score of student :rizky 80

package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Score int
}

type Students []Student

func (s Students) AverageScore() float64 {
	var sum int
	for _, student := range s {
		sum += student.Score
	}
	return float64(sum) / float64(len(s))
}

func (s Students) MinScore() (string, int) {
	var minScore int = math.MaxInt32
	var minName string

	for _, student := range s {
		if student.Score < minScore {
			minScore = student.Score
			minName = student.Name
		}
	}
	return minName, minScore
}

func (s Students) MaxScore() (string, int) {
	var maxScore int = math.MinInt32
	var maxName string

	for _, student := range s {
		if student.Score > maxScore {
			maxScore = student.Score
			maxName = student.Name
		}
	}
	return maxName, maxScore
}

func main() {
	var students Students

	for i := 0; i < 5; i++ {
		var name string
		var score int
		fmt.Printf("Input %d student name: ", i+1)
		fmt.Scan(&name)
		fmt.Printf("Input %d student score: ", i+1)
		fmt.Scan(&score)

		student := Student{
			Name:  name,
			Score: score,
		}
		students = append(students, student)
	}

	averageScore := students.AverageScore()
	minName, minScore := students.MinScore()
	maxName, maxScore := students.MaxScore()

	fmt.Printf("Average score: %.2f\n", averageScore)
	fmt.Printf("Min score of student: %s %d\n", minName, minScore)
	fmt.Printf("Max score of student: %s %d\n", maxName, maxScore)
}

// ! output
// S C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\9_String, Advance Function, Pointer, Struct, Method, Interface, Package & Error Hango run "c:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\9_String, Advance Function, Pointer, Struct, Method, Interface, Package & Error Handling\Praktikum\soal prioritas 1 (80)\program2\main.go"
// Input 1 student name: rizky
// Input 1 student score: 80
// Input 2 student name: kobar
// Input 2 student score: 75
// Input 3 student name: ismail
// Input 3 student score: 70
// Input 4 student name: umam
// Input 4 student score: 75
// Input 5 student name: yopan
// Input 5 student score: 60
// Average score: 72.00
// Min score of student: yopan 60
// Max score of student: rizky 80
// PS C:\Users\R\Documents\GitHub\go-muhammad-reza-hidayat\9_String, Advance Function, Pointer, Struct, Method, Interface, Package & Error Handling>
