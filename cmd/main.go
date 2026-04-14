package main

import (
	"flag"
	"moneytrx/internal/config"
	"moneytrx/internal/repository"
	"moneytrx/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	var pgString string
	flag.StringVar(&pgString, "pgString", "host=localhost port=5432 user=postgres password=postgres dbname=money_trx timezone=UTC", "Postgres connection string")

	conn, err := config.ConnectDb(pgString)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	repo := repository.PgRepo{
		DB: conn,
	}

	ge := gin.Default()
	routes.SetupRoutes(ge, repo)
	err = ge.Run(":8080")
	if err != nil {
		panic(err)
	}
}
