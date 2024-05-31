package user

import (
	"fmt"

	"github.com/steevedevops/simplecruda/src/db"
)

type User struct {
	// bun.BaseModel `bun:"table:users,alias:usr"`
	ID        int    `bun:"id"`
	Username  string `bun:"username" form:"username" binding:"required"`
	Firstname string `bun:"firstname" form:"firstname" binding:"required"`
	Lastname  string `bun:"lastname" form:"lastname"`
	Email     string `bun:"email" form:"email" binding:"required"`
}

func (us *User) FetchUser(search string) ([]User, error) {
	db, ctx, err := db.OpenConnection()
	if err != nil {
		return nil, fmt.Errorf("Nao foi possivel conectar com o banco de dados %v", err)
	}
	defer db.Close()

	// err := db.NewCreateTable().Model((*User)(nil)).Exec(ctx)

	users := []User{}

	if search != "" {

		if err := db.NewSelect().
			Model(&users).
			Where("username = ?", search).
			OrderExpr("id ASC").
			Scan(ctx); err != nil {
			return nil, err
		}
	} else {

		if err := db.NewSelect().
			Model(&users).
			OrderExpr("id ASC").
			Scan(ctx); err != nil {
			return nil, err
		}
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
	db, ctx, err := db.OpenConnection()
	if err != nil {
		return nil, err
		// return nil, fmt.Errorf("Nao foi possivel conectar com o banco de dados %v", err)
	}
	defer db.Close()

	_, err = db.NewInsert().Model(&user).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
