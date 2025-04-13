package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, World!\n")
	}
	http.HandleFunc("/", helloHandler)
	fmt.Println("ok")
	http.ListenAndServe(":8080", nil)

	// log.Println("Starting server on :8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
