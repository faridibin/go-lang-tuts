package main

import "fmt"

func main() {
	//defining a map; long method
	phoneNumbers := make(map[string]string)
	emails := make(map[string]string)

	//assigning values
	phoneNumbers["Farid"] = "+233209092175"
	phoneNumbers["Hamza"] = "+233502066106"

	emails["Farid"] = "faridibin@gmail.com"
	emails["Hamza"] = "infirnohamza@gmail.com"
	emails["noreply"] = "noreply@gmail.com"

	//defining a map; short method
	location := map[string]string{"Farid": "Accra", "Hamza": "Madina"}

	//deleting from map
	delete(emails, "noreply")

	//retrieve all values from map
	fmt.Println(phoneNumbers)
	fmt.Println(emails)
	fmt.Println(location)

	//retrieve value with key from map
	fmt.Println(phoneNumbers["Farid"])
	fmt.Println(emails["Hamza"])

	//get map length
	fmt.Println(len(phoneNumbers))
	fmt.Println(len(emails))
}
