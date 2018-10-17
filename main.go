package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Hello struct {
	Hello string `json:"hello"`
}

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", GetHello).Methods("GET")
	log.Println("starting server....")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetHello(w http.ResponseWriter, r *http.Request) {

	var hello Hello
	hello.Hello = "world"
	json.NewEncoder(w).Encode(hello)
}
