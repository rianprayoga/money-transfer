package routes

import (
	"moneytrx/internal/controller"
	"moneytrx/internal/repository"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, repo repository.PgRepo) {
	c := controller.Controller{
		Db: repo,
	}

	router.POST("/transfer", c.Transfer)
}
