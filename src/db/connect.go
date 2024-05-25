package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/steevedevops/simplecruda/config"
)

func OpenConnection() (*sql.DB, error) {
	conf := config.GetDBConfig()

	connString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", conf.User, conf.Pass, conf.Host, conf.Port, conf.DBname)

	conn, err := sql.Open("postgres", connString)

	if err != nil {
		log.Fatal(err)
	}

	errCheck := conn.Ping()
	return conn, errCheck

}
