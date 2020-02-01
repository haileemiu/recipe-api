package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("main")
	http.HandleFunc("/", fooBarHandler)
	http.ListenAndServe(":8080", nil)
}

func fooBarHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foobar\n")
}

// database name recipe-app
