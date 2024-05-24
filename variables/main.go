package main

import "fmt"

// public decleartion
const LoginToken = "asdfasf345352345dfasfas"

// Variables
func main() {
	//String
	var username string = "Nitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	//bool
	var isLoggIn bool = true
	fmt.Println(isLoggIn)
	fmt.Printf("Variable is of type : %T \n ", isLoggIn)

	//int
	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("Variable is of type : %T \n", smallInt)

	//float
	var floatVal float32 = 255.3434234523452345
	fmt.Println(floatVal)
	fmt.Printf("Variable is of type : %T \n", floatVal)

	var floatVal64 float64 = 255.3434234523452345
	fmt.Println(floatVal64)
	fmt.Printf("Variable is of type : %T \n", floatVal64)

	//default values and alias
	var defaultVal int
	fmt.Println(defaultVal)
	fmt.Printf("Variable is of type : %T \n", defaultVal)

	//default String
	var defaultString string
	fmt.Println(defaultString)
	fmt.Printf("Variable is of type : %T \n", defaultString)

	//Another way of declearing variable
	var name = "Nitesh Paudel"
	fmt.Println(name)
	fmt.Printf("Variable is of type : %T \n", name)

	// dynamic and short type of declearing
	noStudent := 5000
	fmt.Println(noStudent)
	fmt.Printf("Variable is of type : %T \n", noStudent)

	fmt.Println("Login Token : ", LoginToken)
}
