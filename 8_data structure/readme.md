Tipe Data Integer

Tipe data integer terdiri dari dua jenis, yaitu:

1. int: Bilangan cacah (Bilangan positif)
2. uint: Bilangan bulat (bilangan positif maupun negatif)

Penggunaan uint8 dan int8:

- uint8 merepresentasikan angka dengan skala 0 - 255.
- int8 merepresentasikan angka dengan skala -128 - 127.

Ada baiknya menentukan tipe data integer dengan jenis apa yang ingin digunakan untuk menghindari boros memori.

Contoh penggunaan tipe data integer:

```
package main

import "fmt"

func main() {
	var satu uint8 = 89
	var dua int8 = -89

	fmt.Printf("tipe data satu : %T\n", satu)
	fmt.Printf("tipe data dua : %T\n", dua)
}
```

Output program:

```
tipe data satu : uint8
tipe data dua : int8
```

Penggunaan uint untuk angka positif dan penggunaan int8 untuk angka negatif.

Tipe Data String

Tipe data string digunakan untuk merepresentasikan teks atau karakter. Contoh penggunaan tipe data string:

```
package main

import "fmt"

func main() {
	var m string = "reza"

	fmt.Print(m)
}
```

Output program:

```
reza
```

Nil & Zero Value

Nil bukan merupakan tipe data, melainkan sebuah nilai. Variabel yang isi nilainya nil berarti memiliki nilai kosong.

Zero value adalah nilai awal dari sebuah variabel saat belum diberikan nilai secara eksplisit. Berikut adalah zero value dari beberapa tipe data:

- Zero value dari string adalah "" (string kosong).
- Zero value dari bool adalah false.
- Zero value dari tipe numerik non-desimal adalah 0.
- Zero value dari tipe numerik desimal adalah 0.0.

Nil bisa digunakan pada beberapa tipe data seperti pointer, function, slice, map, channel, dan interface kosong atau interface{}.

Sedangkan untuk penggunaan nil pada tipe data selain yang tersebut di atas tidak diperbolehkan.

Contoh penggunaan nil:

```
package main

import "fmt"

func main() {
	var ptr *int = nil
	fmt.Println("Nilai ptr: ", ptr)

	var arr []int = nil
	fmt.Println("Nilai arr: ", arr)

	var mp map[string]int = nil
	fmt.Println("Nilai mp: ", mp)

	var ch chan int = nil
	fmt.Println("Nilai ch: ", ch)

	var iface interface{} = nil
	fmt.Println("Nilai iface: ", iface)
}
```

Output program:

```
Nilai ptr:  <nil>
Nilai arr:  <nil>
Nilai mp:  <nil>
Nilai ch:  <nil>
Nilai iface:  <nil>
```

Dari contoh di atas, variabel ptr, arr, mp, ch, dan iface memiliki nilai kosong atau nil.

Pada Go, penggunaan zero value juga berguna saat membuat variabel dengan tipe data struct atau array yang membutuhkan inisialisasi awal. Dalam hal ini, zero value akan digunakan sebagai nilai awal dari setiap elemen pada struct atau array tersebut.

Contoh penggunaan zero value pada struct dan array:

```
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// membuat variabel struct tanpa inisialisasi
	var p Person
	fmt.Println(p) // outputnya { 0}

	// membuat array dengan panjang tertentu tanpa inisialisasi
	var arr [5]int
	fmt.Println(arr) // outputnya [0 0 0 0 0]
}
```
