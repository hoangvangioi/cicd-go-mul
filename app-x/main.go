package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from App X!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("App X is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
