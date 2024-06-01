package main

import "fmt"

func main() {
	fmt.Println("Loops in go")

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	fmt.Println(days)

	// Method 1
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	fmt.Println("=======================")

	//Method 2
	for i := range days {
		fmt.Println(days[i]) // In this method, it returns only the index
	}
	fmt.Println("=======================")

	//Method 3
	for index, val := range days {
		fmt.Printf("Index : %v , Value : %v \n", index, val)
	}

	//When can also use the go to statement by using goto keyword
}
