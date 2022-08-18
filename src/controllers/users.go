package controllers

import (
	"devbookapi/src/database"
	"devbookapi/src/models"
	"devbookapi/src/repositories"
	"devbookapi/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func InsertUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := user.PrepareUser(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)

	userId, err := userRepository.Insert(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	user.ID = userId

	responses.JSON(w, http.StatusCreated, user)

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
