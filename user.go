package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []User
	db.Find(&users)

	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	var u User

	db.Where("Name = ?", vars["Name"]).First(&u)
	json.NewEncoder(w).Encode(u)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u User
	_ = json.NewDecoder(r.Body).Decode(&u)

	if u.Email != "" && u.Name != "" {
		if err := db.Create(&User{Name: u.Name, Email: u.Email}).Error; err != nil {
			log.Println(err.Error())
			json.NewEncoder(w).Encode(err.Error())

		}
		json.NewEncoder(w).Encode(u)
	} else {
		log.Println("Email and Name are mandatory")
		json.NewEncoder(w).Encode("Email and Name are mandatory")

	}

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	var u User

	if err := db.Where("Name = ?", vars["Name"]).First(&u).Error; err != nil {
		log.Println(err.Error())
		json.NewEncoder(w).Encode(err.Error())

	} else {
		db.Delete(&u)
		json.NewEncoder(w).Encode("User Succesfully deleted")

	}

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email

	db.Save(&user)
	fmt.Fprintf(w, "User Successfully Updated")
}
