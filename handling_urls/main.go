package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://go.dev/doc/effective_go"

func main() {
	fmt.Println("Handling url in go")

	//parsing the url
	res, _ := url.Parse(myUrl)
	fmt.Println("Hostname : ", res.Host)
	fmt.Println("Scheme : ", res.Scheme)
	fmt.Println("Path : ", res.Path)
	fmt.Println("Query : ", res.RawQuery)
	fmt.Println("Port : ", res.Port())

	qParams := res.Query()

	fmt.Println("Formatted Query : ", qParams)

	//creating url
	partsOfUrl := &url.URL{ //we need to pass refrence not the copy
		Scheme: "https",
		Host:   "go.dev",
		Path:   "/doc/effective_go",
	}

	customUrl := partsOfUrl.String()
	fmt.Println("Custom url : ", customUrl)

}
