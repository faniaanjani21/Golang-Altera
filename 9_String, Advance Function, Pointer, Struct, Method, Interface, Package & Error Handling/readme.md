string
string adalah tipe data yang paling umum digunakan di Golang. Tipe data ini digunakan untuk merepresentasikan urutan karakter Unicode yang membentuk sebuah string.

rune adalah tipe data yang merepresentasikan sebuah karakter Unicode, yang terdiri dari satu atau beberapa byte.

byte adalah tipe data yang merepresentasikan sebuah byte dalam urutan byte yang membentuk sebuah string.

[]byte adalah tipe data slice dari byte, yang digunakan untuk merepresentasikan urutan byte yang membentuk sebuah string.

[]rune adalah tipe data slice dari rune, yang digunakan untuk merepresentasikan urutan karakter Unicode yang membentuk sebuah string.

stringer adalah sebuah interface yang digunakan untuk mengubah suatu nilai ke dalam bentuk string, sehingga nilainya dapat dicetak atau ditampilkan.

builder adalah tipe data yang digunakan untuk membangun sebuah string secara dinamis, yaitu dengan menambahkan karakter atau urutan karakter ke dalam string tersebut secara bertahap.

contoh sederhana method

package main

import (
"fmt"
)

// Define a struct named Rectangle
type Rectangle struct {
width float64
height float64
}

// Define a method for Rectangle struct to calculate its area
func (r Rectangle) area() float64 {
return r.width \* r.height
}

func main() {
// Create a new Rectangle object
rect := Rectangle{width: 10, height: 5}

    // Call the area method for the Rectangle object
    area := rect.area()

    // Print the area of the rectangle
    fmt.Printf("The area of the rectangle is: %f\n", area)

}

contoh sederhana struct
package main

import (
"fmt"
)

// Define a struct named Person
type Person struct {
Name string
Age int
Address string
}

func main() {
// Create a new Person object
person := Person{Name: "John", Age: 30, Address: "123 Main St"}

    // Print the person's information
    fmt.Printf("Name: %s, Age: %d, Address: %s\n", person.Name, person.Age, person.Address)

    // Change the person's address
    person.Address = "456 Oak St"

    // Print the updated information
    fmt.Printf("Name: %s, Age: %d, Address: %s\n", person.Name, person.Age, person.Address)

}
