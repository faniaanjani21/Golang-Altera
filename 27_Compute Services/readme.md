Compute Services pada Golang adalah layanan untuk melakukan komputasi pada aplikasi yang dibangun menggunakan bahasa pemrograman Go. Layanan ini menyediakan berbagai fitur seperti pemrosesan data, pengolahan gambar, dan pemrograman paralel. Dengan menggunakan Compute Services pada Golang, aplikasi dapat melakukan komputasi secara efisien dan cepat, sehingga dapat meningkatkan performa dan efektivitas aplikasi. Selain itu, Compute Services pada Golang juga menyediakan dukungan untuk berbagai platform seperti cloud computing dan mobile devices. Dengan demikian, Compute Services pada Golang dapat membantu pengembang untuk membangun aplikasi yang lebih efisien dan efektif.

## contoh program

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano()) // generate random seed
    var nums []int
    for i := 0; i < 10; i++ {
        nums = append(nums, rand.Intn(100)) // generate random numbers from 0 to 99
    }
    fmt.Println("Random numbers:", nums)
    fmt.Println("Sum of numbers:", computeSum(nums))
}

func computeSum(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}
```
