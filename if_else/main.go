package main

import "fmt"

func main() {
	var marks float32
	fmt.Println("If-else in go")
	fmt.Println("Enter your marks : ")
	fmt.Scanf("%f", &marks)
	if marks >= 80 {
		fmt.Println("You are pass")
	} else {
		fmt.Println("your are fail")
	}

	//Another syntax

	var num int
	fmt.Println("Enter any number to check whether it is even or not: ")
	fmt.Scanf("%d", num)

	if val := num; val%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}
}
