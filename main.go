package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "A simple HTTP server")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf(fmt.Sprintf("Number of bytes written: %d", n))
	})
	_ = http.ListenAndServe(":80", nil)
}
