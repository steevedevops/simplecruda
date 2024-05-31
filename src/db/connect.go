package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/steevedevops/simplecruda/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

var Instance *bun.DB

// just bun return
func OpenConnection() (*bun.DB, error) {
	if Instance != nil {
		return Instance, nil
	}
	conf := config.GetDBConfig()
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", conf.User, conf.Pass, conf.Host, conf.Port, conf.DBname)
	sqldbconn, err := sql.Open("postgres", connString)

	if err != nil {
		return nil, nil
	}

	Instance = bun.NewDB(sqldbconn, pgdialect.New())
	Instance.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	errCheck := Instance.Ping()
	return Instance, errCheck
}
