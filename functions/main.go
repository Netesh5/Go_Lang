package main

import "fmt"

func main() {
	fmt.Println("Function in go")

	result := add(5, 7)

	multiResult := multiAdder(5, 6, 7, 8, 9, 10)
	multiResultwithDiffTypeVal, stringVal := multiAdderWithDiffType(5, 6, 7, 8, 9, 10, 11)
	fmt.Println(result)
	fmt.Println(multiResult)
	fmt.Println(multiResultwithDiffTypeVal, stringVal)

}

func add(firstVal int, secondVal int) int {
	return firstVal + secondVal
}

func multiAdder(values ...int) int {
	total := 0
	for _, i := range values {
		total += i
	}
	return total
}

func multiAdderWithDiffType(values ...int) (int, string) {
	total := 0
	for _, i := range values {
		total += i
	}
	return total, "GO is amazing"
}
