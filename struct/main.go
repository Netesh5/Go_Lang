package main

import "fmt"

func main() {
	fmt.Println("Structs in go")

	nitesh := User{"Nitesh", "netesh.paudel@gmail.com", true, 24}
	fmt.Println(nitesh)

	fmt.Printf("Detail of struct : %+v\n", nitesh)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
