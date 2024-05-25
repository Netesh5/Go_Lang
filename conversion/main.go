package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter your age : ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(input)

	newAge, err := strconv.ParseInt(strings.TrimSpace(input), 10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(newAge + 1)

}
