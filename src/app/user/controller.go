package user

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateUserLoginController(db *sql.DB) {
	user := InitUserDB(&User{
		Username:  "steevedevops",
		Firstname: "Steeve",
		Lastname:  "Devops",
		Email:     "steevedev@besoft.com.br",
		IsActive:  true,
	})

	resId, err := user.CreateUser(db)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resId)
}

func BuscarUsuarioController() {
	var user *User
	users, err := user.FetchUser()

	if err != nil {
		log.Fatal(err)
	}
	for user := range users {
		fmt.Println(user)
	}
}
