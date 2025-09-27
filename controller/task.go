package controller

import (
	"awesomeProject/service"
	"awesomeProject/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService *service.TaskService
}

func NewTaskController(s *service.TaskService) *TaskController {
	return &TaskController{taskService: s}
}

func (t *TaskController) CreateTask(ctx *gin.Context) {
	var task *types.CreateTask
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idStr := ctx.GetUint("userID")
	task.UserID = idStr
	err := t.taskService.CreateTask(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}
