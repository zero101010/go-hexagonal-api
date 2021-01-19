package main

import (
	"encoding/json"
	"fmt"

	// "go-api-hexagonal/application/services"
	"go-api-hexagonal/domain"
	"go-api-hexagonal/framework/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Define o user de acordo com a entidade definida em domain
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal("Erro ao receber body")
	}
	database.DB.Create(&user)
	response, err := json.Marshal(user)
	if err != nil {
		log.Fatal("Erro ao adicionar ")
	}
	// services.CreateUser(&user)
	w.Write(response)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var user []domain.User
	database.DB.Find(&user)
	response, err := json.Marshal(user)
	if err != nil {
		log.Fatal("Erro ao adicionar ")
	}
	w.Write(response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	params := mux.Vars(r)
	id := params["id"]
	database.DB.First(&user, id)
	response, err := json.Marshal(user)
	fmt.Printf(string(response))
	if err != nil {
		log.Fatal("Erro ao adicionar ")
	}
	w.Write(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var userDeleted domain.User
	var userFinded domain.User
	params := mux.Vars(r)
	id := params["id"]
	err := database.DB.First(&userFinded, id).Error
	if err != nil {
		panic("Dont find register")
	}
	database.DB.Delete(&userDeleted, id)
	response, err := json.Marshal(userFinded)
	fmt.Printf(string(response))
	if err != nil {
		log.Fatal("Erro ao adicionar ")
	}
	w.Write(response)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	database.ConnectDataBaseSqlight()
	router.HandleFunc("/user", GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/user", CreateUser).Methods("POST")
	fmt.Println(" Servidor Rodando na porta 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
