package controllers

import (
	"net/http"
)

func InsertUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Adicionando usuario"))
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
