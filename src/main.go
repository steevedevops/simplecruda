package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/steevedevops/simplecruda/config"
	"github.com/steevedevops/simplecruda/src/db"
	"github.com/steevedevops/simplecruda/src/db/migrations"
	"github.com/steevedevops/simplecruda/src/routes"
	"github.com/steevedevops/simplecruda/src/shared/utils"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, _, err := db.OpenConnection()
	if err != nil {
		log.Fatal("Não foi possivel fazer a primeira conexao com o banco de dados")
	}
	// defer db.Close()

	app := &cli.App{
		Name: "bun",
		Commands: []*cli.Command{
			utils.NewDBCommand(migrate.NewMigrator(db, migrations.Migrations)),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	routes.Routes(r)
	r.Run(fmt.Sprintf(":%s", config.GetServerPort()))

}
