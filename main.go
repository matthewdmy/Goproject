package main

import (
	"awesomeProject/controller"
	"awesomeProject/dao"
	"awesomeProject/middleware"
	"awesomeProject/pkg/client"
	"awesomeProject/service"

	"github.com/gin-gonic/gin"
)

func main() {
	elasticClient, _ := client.NewEsClient()
	esRepo := dao.NewEsRepo(elasticClient)

	router := gin.Default()
	db := client.DbClient
	userDAO := dao.NewUserDAO(db)
	userService := service.NewUserService(userDAO)
	userController := controller.NewUserController(userService)
	taskDAO := dao.NewTaskDAO(db)
	taskService := service.NewTaskService(taskDAO, esRepo)
	taskController := controller.NewTaskController(taskService)
	u := router.Group("/user")
	{
		u.POST("/register", userController.Register)
		u.POST("/login", userController.LoginUser)
		u.GET("/:id", middleware.JWTAuth(), userController.GetUser)
		u.PUT("/:id", userController.UpdateUser)
		u.DELETE("/:id", userController.DeleteUser)
	}
	t := router.Group("/task")
	{
		t.POST("/tasks", taskController.CreateTask)
		t.GET("/tasks", taskController.GetUserTasks)
		t.DELETE("/tasks/:id", taskController.DeleteTask)
		t.GET("/tasks/search", taskController.SearchTask)
		t.PUT("/tasks/:id", taskController.UpdateTask)
	}
	router.Run(":8080")
}
