package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("."))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle HTTP requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%+v\n", r)
		fmt.Fprintf(w, "Hello, world!")
	})

	// Start server
	fmt.Println("Connect to http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
