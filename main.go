package main

import (
	services "example/service"
	"fmt"
)

func main() {
	var db []*services.User
	userSvc := services.NewUserService(db)

	user1 := userSvc.Register(&services.User{Nama: "Iqbal"})
	fmt.Println(user1)

	user2 := userSvc.Register(&services.User{Nama: "Hamdani"})
	fmt.Println(user2)

	resGet := userSvc.GetUser()
	fmt.Println("---------------- Get all user ----------------")

	for _, v := range resGet {
		fmt.Println(v.Nama)
	}
}
