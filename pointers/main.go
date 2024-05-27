package main

import "fmt"

func main() {
	fmt.Println("Pointers example")

	// var ptr *int

	// fmt.Println("Value of pointer is : ", ptr)

	num := 5
	var ptr *int = &num
	fmt.Println("value of actual pointer is : ", ptr)
	fmt.Println("valur of actual pointer is : ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is : ", *ptr)

}
