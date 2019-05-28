package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	//for the sql
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

//User is
type User struct {
	gorm.Model
	Name  string
	Email string
}

//InitialMigration function to connect database
func InitialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal("Error :", err)
		panic("Failed to connect")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

//AllUsers function
func AllUsers(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "All users Endpoint")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

//NewUsers function
func NewUsers(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "New user Endpoint")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New User Successfully Created")
}

//DeleteUsers function
func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Delete user Endpoint")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "Successfully Deleted User")
}

//UpdateUsers function
func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Update user Endpoint")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email

	db.Save(&user)
	fmt.Fprintf(w, "Successfully Updated User")
}
