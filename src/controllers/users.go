package controllers

import (
	"devbookapi/src/database"
	"devbookapi/src/models"
	"devbookapi/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func InsertUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	userRepository := repositories.NewUserRepository(db)
	userRepository.Insert(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuarios"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuario"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Editando usuario"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Removendo usuario"))
}
