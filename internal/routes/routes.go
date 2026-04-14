package routes

import (
	"moneytrx/internal/controller"
	"moneytrx/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func SetupRoutes(router *gin.Engine, repo repository.PgRepo, redis *redis.Client) {
	c := controller.Controller{
		Db:    repo,
		Redis: redis,
	}

	router.POST("/transfer", c.Transfer)
}
