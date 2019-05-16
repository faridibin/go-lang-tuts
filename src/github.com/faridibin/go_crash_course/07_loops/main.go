package main

import "fmt"

func main() {
	//for loop long method
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//for loop short method
	for i := 0; i < 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	//alternative for loop
	x := 0
	for {
		fmt.Printf("Alternative Number %d\n", x)
		x++

		if x > 10 {
			break
		}
	}

	/**FizzBuzz
	*Loop through 1 - 100;
	*For any number divisible by 3, output "Fizz"
	*For any number divisible by 5, output "Buzz"
	*For any number divisible by 3 and 5, output "FizzBUzz"
	 */
	for i := 1; i < 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
