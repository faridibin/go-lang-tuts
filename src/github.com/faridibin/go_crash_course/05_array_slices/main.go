package main

import "fmt"

func main() {
	//Arrays
	var carArr [3]string

	//Assigning values
	carArr[0], carArr[1], carArr[2] = "Toyota", "Ford", "Honda"

	//shorthand
	carModelArr := [3]string{"Camry", "Fiesta", "Accord"}

	//Slices
	carSlice := []string{"Lexus", "Mustang", "Acura"}

	//shorthand
	carModelSlice := []string{"LFA", "Bullitt", "NSX"}

	fmt.Println(carArr)
	fmt.Println(carModelArr)

	fmt.Println(carSlice)
	fmt.Println(carModelSlice)
}
