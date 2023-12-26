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
		// 根据商品id获取单个商品具体信息
		productGroup.POST("/getProduct", controller.GetProduct)
		// 根据商品id删除商品
		productGroup.POST("/deleteProduct", controller.DeleteProduct)
		// 根据商品id获取数量
		productGroup.POST("/getQuantity", controller.GetProductNum)
	}

	return r
}
