package main

import "fmt"

//Variables in go

func main() {

	//1st method
	var a int
	a = 5
	fmt.Println(a)

	//2nd method
	var b int = 5
	fmt.Println(b)

	//3rd method
	var c = 5
	fmt.Println(c)

	//4th method

	d := 50
	fmt.Println(d)

	fmt.Println(Value)
	fmt.Println(packageVar)
	fmt.Println(Value)

	val := 100
	fmt.Printf("%v,%T", val, val)
}

//Global vairable

var Value int = 500 // To delecre the global vairable just use UpperCase name (pascal case)

//Pacakge varibale

var packageVar int = 5000 // camel case
