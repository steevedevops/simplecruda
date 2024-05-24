package user

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/steevedevops/simplecruda/db"
)

type User struct {
	Id        int64  `json: "id"`
	Username  string `json: "username"`
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
	IsActive  bool   `json: "isactive"`
	Email     string `json: "email"`
}

func InitUserDB(u *User) *User {
	return u
}

func (u *User) CreateUser(db *sql.DB) (int64, error) {

	query := "INSERT INTO users (username, firstname, lastname, is_active, email) VALUES (?, ?, ?, ?, ?);"
	result, err := db.Exec(query, u.Username, u.Firstname, u.Lastname, u.IsActive, u.Email)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (u *User) FetchUser() ([]User, error) {
	var users []User

	rows, err := db.DBConn.Query("Select * from users")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Id, &u.Firstname, &u.Lastname, &u.Email, &u.IsActive, &u.Username); err != nil {
			return nil, fmt.Errorf("fetchUsers %v", err)
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("fetchUsers %v", err)
	}

	return users, nil
}

// func (u *User) FetchUser(db *sql.DB) ([]User, error) {
// 	var users []User

// 	rows, err := db.Query("Select * from users")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var u User
// 		if err := rows.Scan(&u.Id, &u.Firstname, &u.Lastname, &u.Email, &u.IsActive, &u.Username); err != nil {
// 			return nil, fmt.Errorf("fetchUsers %v", err)
// 		}
// 		users = append(users, u)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, fmt.Errorf("fetchUsers %v", err)
// 	}

// 	return users, nil
// }
