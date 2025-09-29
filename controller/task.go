package controller

import (
	"awesomeProject/service"
	"awesomeProject/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService *service.TaskService
}

func NewTaskController(s *service.TaskService) *TaskController {
	return &TaskController{taskService: s}
}

func (t *TaskController) CreateTask(ctx *gin.Context) {
	var task types.CreateTask
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idStr := ctx.GetUint("userID")
	task.UserID = idStr
	err := t.taskService.CreateTask(&task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (t *TaskController) GetUserTasks(ctx *gin.Context) {
	idStr := ctx.GetUint("userID")
	task, err := t.taskService.GetAllTasks(idStr)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (t *TaskController) SearchTask(ctx *gin.Context) {
	query := ctx.Query("query")
	status := ctx.Query("status")

	idStr := ctx.GetUint("userID")
	task, err := t.taskService.GetTask(query, types.TaskStatus(status), idStr)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (t *TaskController) UpdateTask(ctx *gin.Context) {
	var task types.CreateTask
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := ctx.GetUint("userID")
	if err := t.taskService.UpdateTask(uint(id), userID, &task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (t *TaskController) DeleteTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := ctx.GetUint("userID")
	if err := t.taskService.DeleteTask(uint(id), userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
