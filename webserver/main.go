package main

// easy for make web service
// in the next step we can use gin lip for make web service

import (
	"fmt"
	"net/http"
)

func homehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	http.HandleFunc("/", homehandler)
	http.ListenAndServe(":8080", nil)
}
