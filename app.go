package main

import (
	"fmt"

	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "From extension - Hello world! again !!")

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World")

}

func main() {

	http.HandleFunc("/hello", hello)

	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8080", nil)

}
