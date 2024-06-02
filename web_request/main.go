package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://go.dev/"

func main() {
	fmt.Println("Web request in go")
	res, err := http.Get(url)
	checkNilError(err)
	fmt.Printf("Response type :%T\n", res)
	defer res.Body.Close() //Neccessary to close the response

	dataBytes, err := io.ReadAll(res.Body)
	checkNilError(err)
	fmt.Println(string(dataBytes))
}
func checkNilError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
