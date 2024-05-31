package post

import (
	"database/sql"
	"fmt"
	"os/user"
	"time"

	"github.com/steevedevops/simplecruda/src/db"
)

type Post struct {
	ID          int    `bun:", pk,autoincrement"`
	Title       string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
	OwnerID     int64      `yaml:"owner_id"`
	Owner       *user.User `bun:"rel:belongs-to"`
}

func (ps *Post) FetchPost(search string) ([]Post, error) {
	db, _, err := db.OpenConnection()
	if err != nil {
		return nil, fmt.Errorf("Nao foi possivel conectar com o banco de dados %v", err)
	}
	defer db.Close()

	return []Post{}, nil

}

// func (us *Post) FetchUser(search string) ([]Post, error) {
// 	db, ctx, err := db.OpenConnection()
// 	if err != nil {
// 		return nil, fmt.Errorf("Nao foi possivel conectar com o banco de dados %v", err)
// 	}
// 	defer db.Close()
// 	users := []User{}

// 	if search != "" {

// 		if err := db.NewSelect().
// 			Model(&users).
// 			Where("username = ?", search).
// 			OrderExpr("id ASC").
// 			Scan(ctx); err != nil {
// 			return nil, err
// 		}
// 	} else {

// 		if err := db.NewSelect().
// 			Model(&users).
// 			OrderExpr("id ASC").
// 			Scan(ctx); err != nil {
// 			return nil, err
// 		}
// 	}

// 	return users, nil
// }

// func (us *Post) FetchUserById(id int64) (*Post, error) {
// 	// conn, err := db.OpenConnection()
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("Nao foi possivel conectar com o banco de dados %v", err)
// 	// }
// 	// defer conn.Close()

// 	user := &Post{}

// 	// rows, err := conn.QueryRow("SELECT * FROM users limit 1 ")
// 	// if err != nil {
// 	// 	return user, fmt.Errorf("fetchStudents %v", err)
// 	// }
// 	return user, nil
// }

// func (us *Post) Create(user Post) (*User, error) {
// 	db, ctx, err := db.OpenConnection()
// 	if err != nil {
// 		return nil, err
// 		// return nil, fmt.Errorf("Nao foi possivel conectar com o banco de dados %v", err)
// 	}
// 	defer db.Close()

// 	_, err = db.NewInsert().Model(&user).Exec(ctx)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }
