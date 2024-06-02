package main

import "fmt"

func main() {
	fmt.Println("Defer statements in go")
	defer fmt.Println("This is defer statement")
	fmt.Println("Hello, Golang")

	defers()

}

func defers() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i + 1)
	}
}
