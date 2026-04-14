package main

import (
	"flag"
	"moneytrx/internal/model"
	"moneytrx/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type application struct {
	PgString string
	Db       repository.PgRepo
}

func transfer(c *gin.Context) {

	var req model.TrxRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"amount": req.Amount,
	})

}

func main() {
	app := application{}
	flag.StringVar(&app.PgString, "pgString", "host=localhost port=5432 user=postgres password=postgres dbname=money_trx timezone=UTC", "Postgres connection string")

	conn, err := app.connectDb()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	r := gin.Default()

	r.POST("/transfer", transfer)
	r.Run()
}
