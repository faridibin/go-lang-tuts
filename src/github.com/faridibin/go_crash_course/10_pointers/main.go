package main

import "fmt"

func main() {
	//assigning a variabla as pointer
	variable := "Hello!"
	pointer := &variable

	fmt.Println(variable, pointer)

	//pointer data type
	fmt.Printf("%T\n", pointer)

	//use * to read value of variable from pointer
	fmt.Println(*pointer)
	fmt.Println(*&variable)

	//change value of variable with pointer
	*pointer = "Hello Farid!"
	fmt.Println(variable)
}
