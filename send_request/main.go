package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	performPostRequest()
}

func performPostRequest() {
	const myUrl = "http://localhost:8000/post"

	data := url.Values{}

	data.Add("first_name", "Nitesh")
	data.Add("last_name", "Nitesh")
	data.Add("email", "netesh.paudel@gmail.com")
	data.Add("address", "Kathmandu,Nepal")

	res, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	response, _ := io.ReadAll(res.Body)
	fmt.Println(response)

	defer res.Body.Close()
}
