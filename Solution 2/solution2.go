package main

import "fmt"

func main() {

	const (
		Nov = 11 - iota
		Oct
		Sep
	)

	fmt.Println(Nov, Oct, Sep)
}
