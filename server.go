package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port int64 = 8080
)

type Person struct {
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Based     string `json:"university,omitempty"`
	Email     string `json:"email,omitempty"`
}

type Technology struct {
	Name string `json:"name,omitempty"`
}

type Technologies []Technology

// An http.ResponseWriter value assembles the HTTP server's response;
// by writing to it, we send data to the HTTP client.

//An http.Request is a data structure that represents the client HTTP request.

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	//enable cross origin resource sharing
	// good way to make API accesible by JS in-browser client side code
	enableCors(&w)

	w.Header().Set("Content-Type", "application/json")

	tStack := Technologies{
		Technology{Name: "Go Lang"},
		Technology{Name: "Docker"},
		Technology{Name: "Kubernetes"},
		Technology{Name: "AWS"},
		Technology{Name: "Angular 5"},
		Technology{Name: "Node"},
	}

	json.NewEncoder(w).Encode(tStack)
}

func aboutMeRouteHandler(w http.ResponseWriter, r *http.Request) {
	//enable cross origin resource sharing
	// good way to make API accesible by JS in-browser client side code
	enableCors(&w)
	// setting the header to json for the data
	w.Header().Set("Content-Type", "application/json")
	// data
	p := Person{
		FirstName: "Bob",
		LastName:  "Builder",
		Based:     "Builder University",
		Email:     "Bob@builder.com"}
	// encoding data into jason for the frontend
	json.NewEncoder(w).Encode(p)
}
func projectsRouteHandler(w http.ResponseWriter, r *http.Request) {

}
func cvRouteHandler(w http.ResponseWriter, r *http.Request) {

}

func enableCors(w *http.ResponseWriter) {
	// find out what this pointer call means in this
	// i'm thinking the pointer object can call the method that its real object would call
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	// defined router
	router := mux.NewRouter()

	//routes
	router.HandleFunc("/api/", homePageHandler).Methods("GET")
	router.HandleFunc("/api/aboutme", aboutMeRouteHandler).Methods("GET")
	router.HandleFunc("/api/projects", projectsRouteHandler)
	router.HandleFunc("/api/cv", cvRouteHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
