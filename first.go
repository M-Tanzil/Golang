package main
import "fmt"
func main(){
var name string
fmt.Print("Enter your name")
fmt.Scan(&name)
fmt.Println("Hello",name)
 fmt.Print("Enter number 1")
 var a,b int
 fmt.Scan(&a)
 fmt.Print("Enter Second Number ")
 fmt.Scan(&b)
 fmt.Println("Sum is ",a+b)
}