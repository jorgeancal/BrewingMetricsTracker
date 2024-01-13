package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	log.Print("Hello world")
}

func main() {
	http.HandleFunc("/", homeHandler)
	// Starting the server on port 8080
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
