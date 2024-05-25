package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/steevedevops/simplecruda/config"
	"github.com/steevedevops/simplecruda/src/routes"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	// dbServer := db.NewDBserver("webmaster", "pgsql.dev", "localhost", "simplecruda", 5433)
	// err := dbServer.Open()

	// if err != nil {
	// 	panic(err)
	// }
	routes.Routes(r)
	r.Run(fmt.Sprintf(":%s", config.GetServerPort()))

}
