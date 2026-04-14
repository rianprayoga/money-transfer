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
	var redisString string

	flag.StringVar(&pgString, "pgString", "host=localhost port=5432 user=postgres password=postgres dbname=money_trx timezone=UTC", "Postgres connection string")
	flag.StringVar(&redisString, "redisString", "redis://user:password@localhost:6379/0?protocol=3", "redis connection string")

	conn, err := config.ConnectDb(pgString)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	redis := config.ConnectRedis(redisString)

	repo := repository.PgRepo{
		DB: conn,
	}

	ge := gin.Default()
	routes.SetupRoutes(ge, repo, redis)
	err = ge.Run(":8080")
	if err != nil {
		panic(err)
	}
}
