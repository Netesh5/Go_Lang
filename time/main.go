package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time in go")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))

	//creating own time

	createdTime := time.Date(1999, time.November, 26, 02, 55, 0, 0, time.Local)
	fmt.Print(createdTime)

}
