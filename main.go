package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Home creates the server's homepage
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "A simple HTTP server's homepage")
}

// Users shows the users' array at /users
func Users(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("users.json")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "A list of users: ")
	fmt.Fprintln(w, string(content))
}

// func printJSON(json )

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/users", Users)

	_ = http.ListenAndServe(":80", nil)

}
