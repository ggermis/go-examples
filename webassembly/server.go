package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Open localhost:8080/index.html to load the wasm file...")
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}