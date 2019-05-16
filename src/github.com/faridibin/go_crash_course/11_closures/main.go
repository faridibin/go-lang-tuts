package main

import "fmt"

//defing an anonymous function which we use to form a closure
func addition() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := addition()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}
