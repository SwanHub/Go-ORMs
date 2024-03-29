package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

// the model
type User struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() {
	db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password= dbname=test3 sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password= dbname=test3 sslmode=disable")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	// & == 'all'
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password= dbname=test3 sslmode=disable")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New User Successfully created")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password= dbname=test3 sslmode=disable")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)
	fmt.Fprintf(w, "Deleted user")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("postgres", "host=localhost port=5431 user=jacksonprince password= dbname=test3 sslmode=disable")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email
	db.Save(&user)
	fmt.Fprintf(w, "Updated User")
}
