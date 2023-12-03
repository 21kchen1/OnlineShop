package controller

import (
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

func UsersRegister(c *gin.Context) {
	var theUser models.User

	// 自动获取数据
	if err := c.ShouldBind(&theUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg": "c.ShouldBind 获取参数失败",
		})
		return
	}

	if err := service.UsersRegister(&theUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg": "service.UsersRegister 创建失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg": "用户创建成功",
	})
}

func UsersLogin(c *gin.Context) {

}