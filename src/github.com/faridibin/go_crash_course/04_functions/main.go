package main

import "fmt"

func greeting(name string) string {
	return "Hello there" + name
}

func addUp(num1, num2 int) int {
	return num1 + num2
}

//returning multiple values
func values() (int, string) {
	return 5, "Hello World!"
}

//defining variadic functions
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func student(names ...string) {
	fmt.Print(names, " ")
	totalStudents := len(names)
	fmt.Println("Total number of students:", totalStudents)
}

//defining recursive functions
func recursive(num int) int {
	if num == 0 {
		return 1
	}

	return num * recursive(num-1)
}

func main() {
	fmt.Println(greeting("Farid"))
	fmt.Println(addUp(10, 6))

	//retrieve multiple values
	numberVariable, stringVariable := values()

	fmt.Println(numberVariable)
	fmt.Println(stringVariable)

	//retrieve last value
	_, lastVariable := values()
	fmt.Println(lastVariable)

	//variadic functions
	sum(5, 665, 43, 2)

	students := []string{"John Appiah", "Kofi Sintim", "Rebecca Alotey"}
	student(students...)

	//recursive functions
	fmt.Println(recursive(15))
}
