Number (integer)

1. Basic Type: number, string, boolean.
2. Aggregate Type: array dan struct.
3. Reference Type: slice, pointer, map, function, channel
4. Interface Type: interface

- int: Bilangan cacah (Bilangan positif)
- uint: Bilangan bulat (bilangan positif maupun
negatif)

Ada baiknya jika kita menentukan tipe data integer dengan jenis
apa yang ingin kita pakai untuk menghindari boros memori.


pengunaan uint8 dan int8

 uint8 merepresentasikan angka dengan skala 0 - 255 sedangkan
 int8 merepresentasikan angka dengan skala -128 - 127.

	package main

	import "fmt"

	func main() {
		var satu uint8 = 89
		var dua int8 = -89

		fmt.Printf("tipe data satu : %T\n", satu)
		fmt.Printf("tipe data dua : %T\n", dua)

	}

output program 
	tipe data satu : uint8
	tipe data dua : int8

penggunaan uint untuk angka plus dan pengunaan angka int8 untuk mines


string
	package main

	import "fmt"

	func main() {
		var m string = "reza"

		fmt.Print(m)
	}
output
	reza

jadi pada string seperti biasa di format md an nama output utntuk string nya dan di pangill dengan sebutan var nya di mana harus berurutan ya

nil & Zero Value

nil bukan merupakan tipe data, melainkan sebuah nilai. Variabel yang isi nilainya nil berarti memiliki nilai kosong.
	● Zero value dari string adalah "" (string kosong).
	● Zero value dari bool adalah false.
	● Zero value dari tipe numerik non-desimal adalah 0.
	● Zero value dari tipe numerik desimal adalah 0.0
Nil adalah nilai kosong, benar-benar kosong. Nil tidak bisa digunakan pada tipe data
yang sudah dibahas sebelumnya. Ada beberapa tipe data yang bisa di-set nilainya dengan nil, diantaranya:
pointer
	● tipe data function
	● slice
	● map
	● channel
	● interface kosong atau interface{}
