package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File handling in go")
	content := "Hello world, I am learning golang"

	// creating a file using os package
	file, err := os.Create("./file.txt")
	checkNilError(err)

	len, err := io.WriteString(file, content)

	checkNilError(err)
	fmt.Println("Length of file : ", len)

	defer file.Close()
	readFile("./file.txt")
}

func readFile(fileName string) {
	dataByte, err := os.ReadFile(fileName)
	checkNilError(err)

	fmt.Println(string(dataByte))
}

func checkNilError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
