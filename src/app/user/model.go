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
	users := []User{}

	err := db.DBX.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *User) FetchUserById(id int64) (*User, error) {
	user := &User{}

	err := db.DBX.Get(&user, "SELECT * FROM users where id = ? ", id)
	if err != nil {
		return user, fmt.Errorf("fetchStudents %v", err)
	}
	return user, nil
}

func (us *User) CreatUser(user User) (*User, error) {
	query := "insert into users (username, firstname, lastname, email) values (?,?,?,?);"
	result := db.DBX.MustExec(query, user.Username, user.Firstname, user.Lastname, user.Email)
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
		// return nil, fmt.Errorf("addUser Error: %v", err)
	}
	return us.FetchUserById(id)
}
