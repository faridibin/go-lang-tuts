package main

import "fmt"

func greeting(name string) string {
	return "Hello there" + name
}

func addUp(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Farid"))
	fmt.Println(addUp(10, 6))
}
