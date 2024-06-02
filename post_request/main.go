package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Post request in go")
	postRequest()

}

func postRequest() {
	const myUrl = "http://localhost:8000/post"

	//Fake json payload
	reqBody := strings.NewReader(`
{
	"name":"Nitesh Paudel",
	"country":"Nepal",
	"language":"Nepali",
	"code_lang":"Go Lang"
}
`)
	res, err := http.Post(myUrl, "application/json", reqBody)
	if err != nil {
		panic(err)
	}
	content, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
	defer res.Body.Close()
}
