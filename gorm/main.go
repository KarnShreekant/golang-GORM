package main

import (
	"gorm/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//handler function that defines all the API endpoints requests
func handleRequests() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/users", user.AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", user.NewUsers).Methods("POST")
	myRouter.HandleFunc("/user/{name}", user.DeleteUsers).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", user.UpdateUsers).Methods("PUT")
	log.Fatal(http.ListenAndServe("180.151.234.242:8080", myRouter))
}
func main() {
	user.InitialMigration()
	handleRequests()
}
