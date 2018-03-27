package main

import (
	"fmt"
	"log"
	"net/http"
)

// An http.ResponseWriter value assembles the HTTP server's response;
// by writing to it, we send data to the HTTP client.

//An http.Request is a data structure that represents the client HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func main() {
	//  handle all requests to the web root ("/") with handler
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
