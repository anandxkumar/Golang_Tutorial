package main

import "fmt"

func main() {

	fmt.Print("Enter Age: ")

	var age int

	// Taking input from user
	fmt.Scanln(&age)

	if age > 60 {
		fmt.Println("Getting Older")
	} else if age > 30 {
		fmt.Println("Getting Wiser")
	} else if age > 20 {
		fmt.Println("Adulthood")
	} else if age > 10 {
		fmt.Println("Younger")
	} else {
		fmt.Println("Kid")
	}

	fmt.Println(age)
}
