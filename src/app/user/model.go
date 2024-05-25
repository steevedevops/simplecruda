package user

import (
	"fmt"

	"github.com/steevedevops/simplecruda/src/db"
)

type User struct {
	ID        int
	Username  string `form:"username" binding:"required"`
	Firstname string `form:"firstname" binding:"required"`
	Lastname  string `form:"lastname"`
	Email     string `form:"email" binding:"required"`
}

func (us *User) FetchUser(search string) ([]User, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, fmt.Errorf("Nao foi possivel conectar com o banco de dados %v", err)
	}
	defer conn.Close()
	users := []User{}

	rows, err := conn.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user)
		if err != nil {
			fmt.Errorf("Nao foi possivel pegar o dados %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (us *User) FetchUserById(id int64) (*User, error) {
	// conn, err := db.OpenConnection()
	// if err != nil {
	// 	return nil, fmt.Errorf("Nao foi possivel conectar com o banco de dados %v", err)
	// }
	// defer conn.Close()

	user := &User{}

	// rows, err := conn.QueryRow("SELECT * FROM users limit 1 ")
	// if err != nil {
	// 	return user, fmt.Errorf("fetchStudents %v", err)
	// }
	return user, nil
}

func (us *User) CreatUser(user User) (*User, error) {
	// query := "insert into users (username, firstname, lastname, email) values (?,?,?,?);"
	// result := db.DBX.MustExec(query, user.Username, user.Firstname, user.Lastname, user.Email)
	// id, err := result.LastInsertId()
	// if err != nil {
	// 	return nil, err
	// 	// return nil, fmt.Errorf("addUser Error: %v", err)
	// }
	return us.FetchUserById(0)
}
