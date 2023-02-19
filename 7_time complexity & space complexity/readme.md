rangkuman time complexity

Time complexity adalah ukuran seberapa cepat program berjalan dan bagaimana kinerjanya tergantung pada ukuran input yang diberikan. Dalam bahasa Go, kompleksitas waktu dapat dihitung dengan mengukur jumlah operasi elementer yang diperlukan untuk menyelesaikan suatu tugas terhadap ukuran input yang diberikan.

contoh program nya

package main

import "fmt"

func main() {
//Baris keempat membuat variabel "arr" yang merupakan array integer dengan elemen {1, 2, 3, 4, 5}, dan variabel "sum" yang diinisialisasi dengan 0.
    arr := []int{1, 2, 3, 4, 5}
    sum := 0
//Baris kelima menjalankan loop for untuk mengiterasi setiap elemen dalam array "arr", menambahkan nilai setiap elemen ke variabel "sum". Loop ini akan berjalan sebanyak "n" kali, di mana "n" adalah jumlah elemen dalam array "arr".
    for i := 0; i < len(arr); i++ {
        sum += arr[i]
    }

    fmt.Println("Jumlah elemen dalam array:", sum)
}

Time complexity mengukur berapa banyak waktu yang dibutuhkan oleh suatu algoritma untuk menyelesaikan tugas terhadap ukuran input yang diberikan. Biasanya diukur dalam notasi Big O, 

 space complexity mengukur berapa banyak ruang yang dibutuhkan oleh suatu algoritma untuk menyelesaikan tugas terhadap ukuran input yang diberikan.