package main

import "fmt"

const Loginkey = "jiqwdufei"

func main() {
	var username string = "Tanzil"
	fmt.Println(username)
	fmt.Printf("The variable type is : %T \n", username)

	var IsLoggedIn bool = true
	fmt.Println(IsLoggedIn)
	fmt.Printf("The variable type is : %T \n", IsLoggedIn)

	var numberint uint8 = 41
	fmt.Println(numberint)
	fmt.Printf("The variable type is : %T \n", numberint)

	var numberflot float32 = 415.685289456298456
	fmt.Println(numberflot)
	fmt.Printf("The variable type is : %T \n", numberflot)

	// way to assign variable

	var randumnumber = 746882.2
	fmt.Println(randumnumber)
	fmt.Printf("The variable type is : %T \n", randumnumber)

	name := "Tanzil"
	fmt.Println(name)
	fmt.Printf("The variable type is : %T \n", name)

	fmt.Println(Loginkey)
	fmt.Printf("The variable type is : %T \n", Loginkey)
}
