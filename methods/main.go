package main

import "fmt"

func main() {
	fmt.Println("Method in go")
	nitesh := User{"Nitesh", "netesh.paudel@gmail.com", true, 24}
	fmt.Println(nitesh)
	nitesh.getStatus()
	nitesh.assignNewMail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (user User) getStatus() {
	fmt.Println("Is user active : ", user.Status)
}

func (user User) assignNewMail() {
	user.Email = "nitesh.paudel@gmail.com"
	fmt.Println(user.Email)

}
