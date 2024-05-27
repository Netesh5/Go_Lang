package main

import "fmt"

func main() {
	fmt.Println("Array example")

	var items [4]string

	items[0] = "Apple"
	items[1] = "Banana"
	items[3] = "Carrot"

	fmt.Println("Items : ", items)

	var randItems = [3]string{"asdf", "dfasfd", "adfa"}
	fmt.Println("Items : ", randItems)
}
