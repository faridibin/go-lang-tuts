package main

import "fmt"

func main() {
	var name, email = "Farid", "faridibin@gmail.com"
	var age = 100
	const isCool = true

	//shorthand
	hobby := "Coding"
	shirtSize, shoeSize := 14.5, 45

	fmt.Println(name, email, age, isCool)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", hobby)
	fmt.Printf("%T\n", shirtSize)
	fmt.Printf("%T\n", shoeSize)
}
