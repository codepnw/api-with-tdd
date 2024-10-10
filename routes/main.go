package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	app *gin.Engine
)

func InitRoutes(db *gorm.DB) *gin.Engine {
	app = gin.Default()

	app.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	return app
}
