package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DBX *sqlx.DB

type DBserver struct {
	User   string
	Passwd string
	Host   string
	Port   int
	DBname string
}

func NewDBserver(User string, Passwd string, Host string, DBname string, Port int) *DBserver {
	return &DBserver{
		User:   User,
		Passwd: Passwd,
		Host:   Host,
		Port:   Port,
		DBname: DBname,
	}
}

func (ds *DBserver) Open() error {
	conn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", ds.User, ds.Passwd, ds.Host, ds.Port, ds.DBname)

	var err error
	DBX, err = sqlx.Connect("postgres", conn)

	if err != nil {
		log.Fatal(err)
	}

	errCheck := DBX.Ping()
	if errCheck != nil {
		log.Fatal(errCheck)
	}
	log.Printf("DB: `%s` conectado com sucesso", ds.DBname)
	return nil
}
