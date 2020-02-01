package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("main")
	r := mux.NewRouter()
	r.HandleFunc("/recipes", GetRecipesHandler).Methods("GET")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

func GetRecipesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foobar\n")
}

// database name recipe-app
