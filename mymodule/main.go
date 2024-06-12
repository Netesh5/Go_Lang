package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Module in go")
	func1()
	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))

}

func func1() {
	fmt.Println("Mod Init")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to Go lang"))
}
