package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// With local server
func main() {
	getRequest()
}

func getRequest() {
	const myUrl = "http://localhost:8000/get"

	res, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	content, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

	defer res.Body.Close()

	//Another way of handling the get request
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("Bytecount : ", byteCount)
	fmt.Println("Response : ", responseString.String())

}
