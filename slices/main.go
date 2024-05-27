package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in go")

	var items = []string{"Apple", "Ball", "Cat"}

	fmt.Printf("Type of list %T\n:", items)

	fmt.Println("Before append : ", items)

	items = append(items, "Dog")

	fmt.Println("After append : ", items)

	items = append(items, "Elephant", "Fish")

	fmt.Println("After append : ", items)

	items = append(items[1:4])

	fmt.Println("After append : ", items)

	items = append(items[:3])

	fmt.Println("After append : ", items)

	//creating slices using make

	highscores := make([]int, 3)

	highscores[0] = 123
	highscores[1] = 213
	highscores[2] = 321

	fmt.Println(highscores)

	highscores = append(highscores, 232, 233)

	fmt.Println(highscores)

	sort.Ints(highscores)

	fmt.Println(highscores)

	fmt.Println(sort.IntsAreSorted(highscores))

}
