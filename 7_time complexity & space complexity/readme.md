Time complexity adalah ukuran seberapa cepat program berjalan dan bagaimana kinerjanya tergantung pada ukuran input yang diberikan. Dalam bahasa Go, kompleksitas waktu dapat dihitung dengan mengukur jumlah operasi elementer yang diperlukan untuk menyelesaikan suatu tugas terhadap ukuran input yang diberikan. Time complexity diukur dalam notasi Big O.

Contoh program Go di bawah ini menghitung jumlah elemen dalam array dan memiliki time complexity O(n), di mana n adalah jumlah elemen dalam array tersebut:

```go
package main

import "fmt"

func main() {
    arr := []int{1, 2, 3, 4, 5}
    sum := 0

    for i := 0; i < len(arr); i++ {
        sum += arr[i]
    }

    fmt.Println("Jumlah elemen dalam array:", sum)
}
```

Space complexity mengukur berapa banyak ruang yang dibutuhkan oleh suatu algoritma untuk menyelesaikan tugas terhadap ukuran input yang diberikan. Space complexity juga diukur dalam notasi Big O.
