package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Website"
	fmt.Println(welcome)

	fmt.Println("Please give rate the website")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	fmt.Printf("The type of rating is %T", input)
}
