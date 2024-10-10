package routes

import (
	"api-with-tdd/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	app *gin.Engine
)

func InitRoutes(db *gorm.DB) *gin.Engine {
	app = gin.Default()
	handler := handler.NewHandler(db)

	app.GET("/healthcheck", handler.HealthCheck)

	app.POST("/tasks", handler.CreateTask)
	app.GET("/tasks", handler.GetTasks)
	app.DELETE("/tasks/:id", handler.DeleteTask)

	return app
}
