package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps Example")

	langs := make(map[string]string)

	langs["drt"] = "dart"
	langs["js"] = "javascript"
	langs["py"] = "python"
	langs["rb"] = "ruby"

	fmt.Println("List of all langs in map: ", langs)

	fmt.Println("Js : ", langs["js"])

	//delete item form maps

	delete(langs, "rb")

	fmt.Println("List of all langs in map: ", langs)

	//loops in maps

	for key, value := range langs {
		fmt.Printf("Key %v : Value %v \n", key, value)
	}

}
