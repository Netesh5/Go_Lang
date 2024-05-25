package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// default way of taking input
	var name string
	fmt.Printf("Enter your name : ")
	fmt.Scan(&name)
	fmt.Println(name)

	//Bufio Method
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your name : ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(input)
}
