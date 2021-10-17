package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home Bage")
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> About Page </h1>")
}
func main() {
	// routes
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	fmt.Println("listening to the server")
	// listen to server
	http.ListenAndServe(":3000", nil)
}
