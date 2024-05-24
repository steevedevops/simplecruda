package db

import (
	_ "github.com/lib/pq" // Importar o driver PostgreSQL
)

// var DBConn *sql.DB

// cannot define new methods on non-local type sql.DBcompilerInvalidRecv

// type DBserver struct {
// 	User   string
// 	Passwd string
// 	Host   string
// 	Port   int
// 	DBname string
// }

// func NewDBserver(User string, Passwd string, Host string, DBname string, Port int) *DBserver {
// 	return &DBserver{
// 		User:   User,
// 		Passwd: Passwd,
// 		Host:   Host,
// 		Port:   Port,
// 		DBname: DBname,
// 	}
// }

// func (ds *DBserver) ConnectDB() (*sql.DB, error) {
// 	conn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", ds.User, ds.Passwd, ds.Host, ds.Port, ds.DBname)
// 	db, err := sql.Open("postgres", conn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

//		errCheck := db.Ping()
//		if errCheck != nil {
//			log.Fatal(errCheck)
//		}
//		log.Printf("DB: `%s` conectado com sucesso", ds.DBname)
//		return db, nil
//	}
// func InitDB(User string, Passwd string, Host string, DBname string, Port int) error {
// 	conn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", User, Passwd, Host, Port, DBname)
// 	DBConn, err := sql.Open("postgres", conn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	errCheck := DBConn.Ping()
// 	if errCheck != nil {
// 		log.Fatal(errCheck)
// 	}
// 	log.Printf("DB: `%s` conectado com sucesso", DBname)
// 	return nil
// }
