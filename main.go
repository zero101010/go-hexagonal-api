package main

import (
	"fmt"
	"go-api-hexagonal/application/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// func init() {
// 	database.ConnectDataBaseMysql()
// }
func main() {
	router := mux.NewRouter().StrictSlash(true)

	// User Routes
	router.HandleFunc("/user", services.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", services.GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", services.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user", services.CreateUser).Methods("POST")
	// Quotes Routes
	fmt.Println(" Servidor Rodando na porta 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
