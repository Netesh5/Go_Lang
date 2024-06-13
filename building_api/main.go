package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Creating the model

type Course struct {
	Id         string  `json:"id"`
	CourseName string  `json:"course_name"`
	Price      string  `json:"price"`
	Auther     *Auther `json:"author"`
}

type Auther struct {
	FullName string `json:"full_name"`
	Website  string `json:"website"`
}

//creating slices to store the value

var courses []Course

//creating the middelware

func (c *Course) isEmpty() bool {
	return c.Id == "" && c.CourseName == ""
}

//controller

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to go api</h2>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Gell all the courses")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

func getOneCourese(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, book := range courses {
		if book.Id == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")
	return

}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create course")
	//checking empyt data
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//checking { }data

	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data inside")
		return
	}
	seed := time.Now().UnixNano()
	src := rand.NewSource(seed)
	rnd := rand.New(src)
	num := rnd.Intn(100)
	course.Id = strconv.Itoa(num)
	courses = append(courses, course)
	json.NewEncoder(w).Encode(courses)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	params:=mux.Vars(r)

	for index ,course:range courses{
		if course.id == params["id"] {
		courses=append(courses[:index],courses[index+1:]... )
		var course Course

		json.NewDecoder(r.Body).Decode(&course)
		course.id=params["id"]
		courses=append(db,course)
		json.NewEncoder(w).Encoder(course)
		}
		
	}
}
func main() {

}
