package main

import "fmt"

func main() {
	var a int
	c := 0
	fmt.Println("Enter the Prime Number : ")
	fmt.Scan(&a)
	if a <= 1 {
		fmt.Println("It not a prime")
		return
	} else {
		for i := 2; i*i <= a; i++ {
			if a%i == 0 {
				c++
				break
			}
		}
		if c == 0 {
			fmt.Println("its a prime")
		} else {
			fmt.Println("no")
		}
	}
}
