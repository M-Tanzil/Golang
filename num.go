package main
import "fmt"
func main(){
  for i :=0;i<=100;i++{
  fmt.Println(i)}
  var n,b int
  fmt.Println("Enter the number you want sum")
  fmt.Scan(&n)
  for i:=0;i<=n;i++{
    b=b+i
  }
  fmt.Println(b)
}