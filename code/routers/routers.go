package routers

import (
	"github.com/gin-gonic/gin"
	"onlineshop/controller"
)

/**
 * @File : routers.go
 * @Description : 设置路由
 * @Author : chen
 * @Date : 2023/12/03
 */

func SetupRouters() *gin.Engine {
	r := gin.Default()

	// 路由组 具体待定
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", controller.UsersLogin) //路由改为post
		userGroup.POST("/register", controller.UsersRegister)
	}

	return r
}
