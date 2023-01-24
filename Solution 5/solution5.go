package main

import "fmt"

func multiple(a, b, c int) int {
	total := a * b * c
	return total
}
func main() {

	const (
		Nov = 11 - iota
		Oct
		Sep
	)

	fmt.Print("Enter number a, b, c: ")
	var a int
	var b int
	var c int

	fmt.Scanln(&a, &b, &c)

	fmt.Println(multiple(a, b, c))
}
