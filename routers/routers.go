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

	// 用户相关路由
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", controller.UsersLogin)
		userGroup.POST("/register", controller.UsersRegister)
	}

	// 商品相关路由
	productGroup := r.Group("/product")
	{
		productGroup.POST("/getList", controller.GetProductList)
		// 添加其他商品相关的路由，例如添加商品、更新商品信息等
	}

	return r
}
