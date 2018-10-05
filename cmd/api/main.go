package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/translate", translateFromEnglish).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func translateFromEnglish(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
