package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// jsonPrettyPrint allows printing JSON as separate lines of text
func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

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
	myString := string(content[:])
	fmt.Fprintln(w, jsonPrettyPrint(myString))
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/users", Users)

	_ = http.ListenAndServe(":80", nil)

}
