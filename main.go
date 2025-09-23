package main

import (
	"awesomeProject/middleware"
	"awesomeProject/models"
	"awesomeProject/pkg/client"
	"fmt"
	"net/http"
	"strconv"

	"sync"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var (
	users   = make(map[int]User)
	nextID  = 1
	userMux sync.Mutex
)

func main() {
	router := gin.Default()
	u := router.Group("/user")
	{
		u.POST("/register", RegisterUser)
		u.POST("/login", LoginUser)
		u.GET("/:id", middleware.JWTAuth(), GetUser)
		u.PUT("/:id", UpdateUser)
		u.DELETE("/:id", DeleteUser)
	}

	router.Run(":8080")
}

// 注册
func RegisterUser(c *gin.Context) {
	var req User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := client.DbClient

	err := db.Create(&models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

// 登录
func LoginUser(c *gin.Context) {
	var req User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Username == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty username or password"})
		return
	}
	db := client.DbClient
	var user models.User
	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	if user.Password != req.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}

	token, err := middleware.GenerateToken(uint(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "user": user, "token": token})
	//for _, user := range users {
	//	if user.Username == req.Username && user.Password == req.Password {
	//		c.JSON(http.StatusOK, gin.H{"message": "登录成功", "user": user})
	//		return
	//	}
	//}
	//
	//c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
}

// 获取用户信息
func GetUser(c *gin.Context) {
	idStr := c.GetUint("userID")
	fmt.Println(idStr)
	//id, err := strconv.Atoi(idStr.s)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	db := client.DbClient
	var user models.User

	if err := db.First(&user, idStr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": user})
}

// 更新用户信息
func UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := client.DbClient
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&user).Updates(req).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新成功", "user": user})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	// id exist?
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := client.DbClient
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功", "user": user})
}
