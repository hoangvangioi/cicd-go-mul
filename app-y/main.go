package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from App Y!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("App Y is running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
