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
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	
	// sending a message as a http response to the client through the ResponseWriter, using Fprintf
	fmt.Fprintf(w, "Welcome to the URL shortener service !")
}

