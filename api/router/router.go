package router

import "github.com/gorilla/mux"

func New() *mux.Router {

	r := mux.NewRouter()

	//Handler to Create a record
	r.HandleFunc("/users", CreateUserHandler).Methods("POST")

	//Handler to Get a record by ID
	r.HandleFunc("/user/{id}", GetUserHandler).Methods("GET")

	//Handler to get all the records
	r.HandleFunc("/user", GetAllUserHandler).Methods("GET")

	//Handler to delete a record by ID
	r.HandleFunc("/user/{id}", DeleteUserHandler).Methods("DELETE")

	//Handler to update a record by ID
	r.HandleFunc("/user/{id}", UpdateUserHandler).Methods("PUT")
	return r
}
