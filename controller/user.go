package controller

import (
	"fmt"
	"net/http"
	"onlineshop/models"
	"onlineshop/service"

	"github.com/gin-gonic/gin"
)

/**
 * @File : user.go
 * @Description : 用户相关的控制器
 * @Author : chen
 * @Date : 2023/12/03
 */

// 注册模块
func UsersRegister(c *gin.Context) {
	var theUser models.User

	// 自动获取数据
	if err := c.ShouldBind(&theUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "c.ShouldBind 获取参数失败",
		})
		return
	}

	if err := service.CheckUserExistByName(&theUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "CheckUserExistByName 用户名重复",
		})
		return
	}

	if err := service.CreateUsers(&theUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "service.UsersRegister 创建失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "UsersRegister 用户创建成功",
	})
}

// 用户登录模块
func UsersLogin(c *gin.Context) {
	fmt.Println("Entering UsersLogin")

	var loginData models.User

	// 手动获取参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	loginData.UserName = username
	loginData.PassWord = password
	fmt.Println("UsersLogin - Username:", username, "Password:", password)
	// 调用验证函数
	if err := service.CheckUserLogin(&loginData); err != nil {
		fmt.Println("Error in CheckUserLogin:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       err.Error(),
		})
		return
	}

	// 获取用户信息
	user, err := models.GetUserByName(loginData.UserName)
	if err != nil {
		fmt.Println("Error in GetUserByName:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "获取用户信息失败",
		})
		return
	}

	// 验证成功，返回相应信息
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "UsersLogin 用户登录成功",
		"userId":    user.ID,
	})
}

func UserGetInf(c *gin.Context) {
	fmt.Println("Go To UserGetInf")
	var requestData struct {
		UserId   int    `json:"userid"`
		UserName string `json:"username" binding:"required"`
		PassWord string `json:"password" binding:"required"`
		PhoneNum string `json:"phoneNum"`
		UserType int    `json:"userType"`
		Email    string `json:"email" binding:"required"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "获取参数失败",
		})
		return
	}

	// 获取商品
	user, err := models.GetUserByID(requestData.UserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "UserGetInf失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "获取用户成功",
		"data":      user,
	})

}
