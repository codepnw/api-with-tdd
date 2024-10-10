package handler

import (
	"api-with-tdd/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *handler {
	return &handler{db: db}
}

func (h *handler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func (h *handler) CreateTask(ctx *gin.Context) {
	var newTask entity.Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed create task"})
		return
	}

	h.db.Create(&newTask)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success create task",
	})
}

func (h *handler) GetTasks(ctx *gin.Context) {
	tasks := []entity.Task{}

	h.db.Find(&tasks)

	ctx.JSON(http.StatusOK, tasks)
}

func (h *handler) DeleteTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "id invalid",
		})
		return
	}

	h.db.Delete(&entity.Task{}, id)

	ctx.JSON(http.StatusOK, gin.H{"message": "success delete task"})
}
