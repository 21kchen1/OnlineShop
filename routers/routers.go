package routers

import (
	"onlineshop/controller"

	"github.com/gin-gonic/gin"
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
		// 添加商品
		productGroup.POST("/addProduct", controller.AddProduct)
		// 根据商品id获取单个商品具体信息
		productGroup.POST("/getProduct", controller.GetProduct)
		// 根据商品id删除商品
		productGroup.POST("/deleteProduct", controller.DeleteProduct)
		// 根据商品id修改商品信息
		productGroup.POST("/editProduct", controller.EditProduct)
		// 根据商品id获取数量
		productGroup.POST("/getQuantity", controller.GetProductNum)
		// 根据商品id修改库存数量
		productGroup.POST("/editQuantity", controller.EditProductNum)
	}

	// 收藏夹列表
	favoritesGroup := r.Group("/favorites")
	{
		// 根据 id 获取收藏夹列表
		favoritesGroup.POST("/getList", controller.GetFavoritesList)
		// 创建收藏夹
		favoritesGroup.POST("/add", controller.AddFavorites)
		// 删除收藏夹
		favoritesGroup.POST("/delete", controller.DeleteFavorites)
		// 修改收藏夹名称
		favoritesGroup.POST("/updata", controller.UpdataFavoName)
		// 获取收藏夹中商品id 列表
		favoritesGroup.POST("/getProduct", controller.GetFavoProductList)
		// 收藏夹添加商品
		favoritesGroup.POST("/addProduct", controller.AddFavoProduct)
		// 删除收藏夹物品
		favoritesGroup.POST("/deleteProduct", controller.DeleteFavoProduct)
	}
	// test
	return r
}
