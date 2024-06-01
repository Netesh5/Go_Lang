package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch cases in go")
	diceNum := rand.Intn(6) + 1
	fmt.Printf("Dice Number : %v\n", diceNum)
	switch diceNum {
	case 1:
		fmt.Println("Move one position forward.")
	case 2:
		fmt.Println("Move two position forward.")
	case 3:
		fmt.Println("Move three position forward.")
	case 4:
		fmt.Println("Move four position forward.")
	case 5:
		fmt.Println("Move five position forward.")
	case 6:
		fmt.Println("Move six position forward and Reroll the dice.")
		main()
	}

}
