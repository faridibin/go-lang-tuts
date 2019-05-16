package main

import "fmt"

func main() {
	//range with slices
	students := []string{"Ntim", "Cole", "Adjoa", "Lois", "Mariam", "Kofi", "Paul", "Owusu", "Edward", "John"}

	//loop through students with index
	for i, student := range students {
		fmt.Printf("Position %d - "+student+"\n", i)
	}

	//loop through students without index
	for _, student := range students {
		fmt.Printf("Student - " + student + "\n")
	}

	//use range to sum up numbers
	numbers := []int{4, 54, 343, 2, 325, 76, 87, 8, 5, 998, 43}
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	fmt.Println("Sum is", sum)

	//range with maps
	electricCars := map[string]string{"Tesla": "Model X", "Mercedes": "EQC 2020", "Audi": "e-tron", "Volvo": "Polestar"}

	//keys and values
	for car, model := range electricCars {
		fmt.Printf("%s made the %s\n", car, model)
	}
}
