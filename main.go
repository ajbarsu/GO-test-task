package main

import (
	"fmt"
	"net/http"
)

// Home creates the server's homepage
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "A simple HTTP server's homepage")
}

// Users shows the users' array at /users
func Users(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "A list of users")
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/users", Users)

	_ = http.ListenAndServe(":80", nil)

}
