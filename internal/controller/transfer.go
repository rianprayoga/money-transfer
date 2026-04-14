package controller

import (
	"moneytrx/internal/model"
	"moneytrx/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Db repository.PgRepo
}

func (ct *Controller) Transfer(c *gin.Context) {

	// ct.Db.ReduceBalance()

	var req model.TrxRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"amount": req.Amount,
	})
}
