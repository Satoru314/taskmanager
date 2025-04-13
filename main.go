package main

import (
	"io"
	"net/http"
	"log"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, World!\n")
	}
	http.HandleFunc("/", helloHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
