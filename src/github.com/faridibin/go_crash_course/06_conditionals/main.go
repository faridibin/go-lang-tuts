package main

import "fmt"

func main() {
	//if else
	x := 100
	y := 34

	if x >= y {
		fmt.Printf("%d is greater than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than or equal to %d\n", y, x)
	}

	//else if
	doorStatus := "open"

	if doorStatus == "open" {
		fmt.Println("The front door is open")
	} else if doorStatus == "closed" {
		fmt.Println("The front door is closed")
	} else {
		fmt.Println("The front door either open or closed")
	}

	//switch case
	color := "green"

	switch color {
	case "red":
		fmt.Println("Color is red")
	case "yellow":
		fmt.Println("Color is yellow")
	case "green":
		fmt.Println("Color is green")
	default:
		fmt.Println("Color is " + color)
	}
}
