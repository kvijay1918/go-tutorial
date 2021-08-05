package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"vijay/project-1/api/db"
	"vijay/project-1/api/model"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println("Error in decoding request", err)
		log.Println("Unable to parse the request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var dbConnector db.User
	dbConnector = db.NewUserDao()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		log.Println("Error in hashing password")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Password = string(passwordHash)

	data, err := dbConnector.Insert(&user)
	if err != nil {
		log.Println("Unable to create an user")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.Password = ""
	jsonData, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("VJ 1")
	id := mux.Vars(r)
	log.Println("map string : ", id)
	userId, _ := strconv.Atoi(id["id"])

	log.Println("VJ 2")
	fmt.Println("Vijay user id : ", userId)

	var dbConnector db.User
	dbConnector = db.NewUserDao()

	log.Println("VJ 3")
	data, err := dbConnector.Get(userId)
	if err != nil {
		log.Println("Unable to get data")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.Password = ""
	jsonData, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	var dbConnector db.User
	dbConnector = db.NewUserDao()

	log.Println("VIJAY")
	data, err := dbConnector.GetAll()
	if err != nil {
		log.Println("Unable to retrieve all the records")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		log.Println("Unable to parse the request message")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)
	userId, _ := strconv.Atoi(id["Id"])
	user.Id = userId

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		log.Println("Unable to Hash the password")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Password = string(passwordHash)

	var dbConnector db.User
	dbConnector = db.NewUserDao()

	data, err := dbConnector.Update(&user)
	if err != nil {
		log.Println("Unable to update the user")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	userId, _ := strconv.Atoi(id["Id"])

	var dbConnector db.User
	dbConnector = db.NewUserDao()

	deletedID, err := dbConnector.Delete(userId)
	if err != nil {
		log.Println("Unable to delete the record")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(deletedID)))

}
