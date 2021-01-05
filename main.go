package main

import (
	"encoding/json"
	"go-api-hexagonal/domain"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	user := domain.User{"1", "Igor", 24}
	response, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(response)
}

func getQuote(w http.ResponseWriter, r *http.Request) {
	user := domain.User{"1", "Igor", 24}
	quote := domain.Quote{"1", "ITSA4", 12.2, 3, &user}
	response, err := json.Marshal(quote)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(response)
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/user", getUser).Methods("GET")
	router.HandleFunc("/quote", getQuote).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
