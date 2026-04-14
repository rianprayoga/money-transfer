package controller

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"moneytrx/internal/model"
	"moneytrx/internal/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Controller struct {
	Db    repository.PgRepo
	Redis *redis.Client
}

func (ct *Controller) Transfer(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var req model.TrxRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	trxRecord, err := ct.Db.ReduceBalance(ctx, 1, req.Amount)
	if err != nil {
		if errors.Is(err, repository.ErrInsufucientBalance) || errors.Is(err, repository.ErrMerchantNotFound) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	trxRecord.Success = req.Success

	v, _ := json.Marshal(trxRecord)
	err = ct.Redis.Publish(ctx, "trx", v).Err()
	if err != nil {
		log.Println(err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"amount": req.Amount,
	})
}
