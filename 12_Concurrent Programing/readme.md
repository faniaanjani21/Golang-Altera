current programing

program squental
merupakan program yang di exsekusi secara brurutan

for i:=1 to 10 do
fmt.Println(i)

program paralel
merupakan program yang di exsekusi secara bersamaan

go func(){
for i:=1 to 10 do
fmt.Println(i)
}()

program concurrent
merupakan program yang di exsekusi secara bersamaan dan berurutan

func main(){
go func(){
for i:=1 to 10 do
fmt.Println(i)
}()
fmt.Scanln()
}
