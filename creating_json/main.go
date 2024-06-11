package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Creating Json")
	encodeJson()
}

type course struct {
	Name     string `json:"course_name"` //alias , if we use this kind of syntax then we can change the name in respone
	Price    float64
	Platform string
	Password string   `json:"-"`              // for removing the field
	Tags     []string `json:"tags,omitempty"` //omitempty removes the field whose value is nil
}

func encodeJson() {
	courses := []course{
		{"Go Lang", 355.99, "youtube", "test@123", []string{"Backend dev", "Go"}},
		{"Go Lang", 355.99, "youtube", "test@1234", []string{"Backend dev", "Go"}},
		{"Go Lang", 355.99, "youtube", "test@12345", []string{"Backend dev", "Go"}},
		{"Go Lang", 355.99, "youtube", "test@123456", nil},
	}

	//encoding

	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
