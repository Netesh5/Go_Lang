package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Consume Json")
	decodeJson()

}

type course struct {
	Name     string `json:"course_name"` //alias , if we use this kind of syntax then we can change the name in respone
	Price    float64
	Platform string
	Password string   `json:"-"`              // for removing the field
	Tags     []string `json:"tags,omitempty"` //omitempty removes the field whose value is nil
}

func decodeJson() {
	jsonData := []byte(`
	{
	
                "course_name": "Go Lang",
                "Price": 355.99,
                "Platform": "youtube",
                "tags": [
                        "Backend dev",
                        "Go"
                ]
    
	}
				`)

	var courseData course
	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonData, &courseData)
		fmt.Printf("%#v\n", courseData) // We use %#v synatx to print intrefaces
	} else {
		fmt.Println("Json is not valid ")
	}
}
