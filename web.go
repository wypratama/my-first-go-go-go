package main

import (
	"encoding/json"
	services "example/service"
	"fmt"
	"net/http"
)

var PORT = ":8080"
var db []*services.User
var userDb = services.NewUserService(db)

func main() {

	http.HandleFunc("/", greet)

	http.HandleFunc("/user", getUser)

	http.HandleFunc("/register", createUser)

	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "go web server application"
	fmt.Fprint(w, msg)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		usr := userDb.GetUser()

		json.NewEncoder(w).Encode(usr)
	}

	// http.Error(w, "Invalid Method", http.StatusBadRequest) ==> dapet error superflouse response....
}

func createUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nama := r.FormValue("nama")

		userDb.Register(&services.User{Nama: nama})
		fmt.Fprint(w, nama+" berhasil didaftarkan")
	}

}
