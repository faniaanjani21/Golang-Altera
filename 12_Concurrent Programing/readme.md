Program Sequential
Merupakan program yang dieksekusi secara berurutan. Contohnya sebagai berikut:

```
for i:=1 to 10 do
fmt.Println(i)
```

Program di atas akan menampilkan angka 1 hingga 10 secara berurutan.

Program Parallel
Merupakan program yang dieksekusi secara bersamaan. Contohnya sebagai berikut:

```
go func(){
for i:=1 to 10 do
fmt.Println(i)
}()
```

Program di atas akan menampilkan angka 1 hingga 10 secara bersamaan.

Program Concurrent
Merupakan program yang dieksekusi secara bersamaan dan berurutan. Contohnya sebagai berikut:

```
func main(){
go func(){
for i:=1 to 10 do
fmt.Println(i)
}()
fmt.Scanln()
}
```

Program di atas akan menampilkan angka 1 hingga 10 secara bersamaan dan akan menunggu input dari pengguna sebelum program selesai dieksekusi.

Untuk membuat cleancode, sebaiknya kita menggunakan format yang sama untuk setiap contoh program dan memberikan penjelasan yang jelas dan singkat untuk setiap jenis program. Selain itu, sebaiknya kode program ditulis dengan rapi dan mudah dibaca. Contoh cleancode untuk program sequential adalah sebagai berikut:

```
// Program Sequential
// Merupakan program yang dieksekusi secara berurutan

for i := 1; i <= 10; i++ {
    fmt.Println(i)
}
```

Contoh cleancode untuk program parallel adalah sebagai berikut:

```
// Program Parallel
// Merupakan program yang dieksekusi secara bersamaan

go func() {
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}()
```

Contoh cleancode untuk program concurrent adalah sebagai berikut:

```
// Program Concurrent
// Merupakan program yang dieksekusi secara bersamaan dan berurutan

func main() {
    go func() {
        for i := 1; i <= 10; i++ {
            fmt.Println(i)
        }
    }()
    fmt.Scanln()
}
```
