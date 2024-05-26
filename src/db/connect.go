package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/steevedevops/simplecruda/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

// just bun return
func OpenConnection() (*bun.DB, context.Context, error) {
	conf := config.GetDBConfig()

	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", conf.User, conf.Pass, conf.Host, conf.Port, conf.DBname)

	sqldbconn, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, nil, err
	}

	db := bun.NewDB(sqldbconn, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	errCheck := db.Ping()
	ctx := context.Background()
	return db, ctx, errCheck
}

// func OpenConnection() (*sql.DB, *bun.DB, error) {
// 	conf := config.GetDBConfig()

// 	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", conf.User, conf.Pass, conf.Host, conf.Port, conf.DBname)

// 	sqldbconn, err := sql.Open("postgres", connString)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	db := bun.NewDB(sqldbconn, pgdialect.New())

// 	// errCheck := sqldbconn.Ping()
// 	errCheck := db.Ping()
// 	return sqldbconn, db, errCheck

// }
