package main

import (
	"github.com/gin-gonic/gin"
	"github.com/steevedevops/simplecruda/src/db"
	"github.com/steevedevops/simplecruda/src/routes"
)

func main() {
	r := gin.Default()
	dbServer := db.NewDBserver("webmaster", "pgsql.dev", "localhost", "simplecruda", 5433)
	err := dbServer.Open()

	if err != nil {
		panic(err)
	}

	routes.Routes(r)
	r.Run(":8000")

}
