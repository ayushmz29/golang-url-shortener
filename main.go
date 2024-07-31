package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = ":8080" // port to run the server on
)

func main() {
	http.HandleFunc("/", handleRoot)

	fmt.Println("Starting server on", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	// why Fprintf instead of Println
	fmt.Fprintf(w, "Welcome to the URL shortener service !")
}
// what should be the branch name for this commit, if Im planning to have front end and backend part?

