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
		//用户登录
		userGroup.POST("/login", controller.UsersLogin)
		//用户注册
		userGroup.POST("/register", controller.UsersRegister)
		//管理员更新用户信息
		userGroup.POST("/update", controller.UpdateUserInfo)
		//管理员新增用户
		userGroup.POST("/addUser", controller.AddUser)
		//管理员删除用户
		userGroup.POST("/delete", controller.DeleteUser)
		//管理员获取用户列表
		userGroup.GET("/getList", controller.GetUserList)

	}

	//商家管理路由
	sellerGroup := r.Group("/seller")
	{
		//获取商家列表
		sellerGroup.GET("/getList", controller.GetSellerList)
		//增加商家
		sellerGroup.POST("/addSeller", controller.AddSeller)
		//删除商家
		sellerGroup.POST("/delete", controller.DeleteSeller)
		//修改商家信息
		sellerGroup.POST("/updata", controller.UpdateSellerInfo)
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
		// 根据商品id查询其所有评论
		productGroup.POST("/getComment", controller.GetCommentsByProductID)
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

	// 评论相关路由
	commentGroup := r.Group("/comment")
	{
		// 用户发表评论
		commentGroup.POST("/add", controller.AddComment)
		// 用户回复评论
		commentGroup.POST("/reply", controller.AddReply)
	}
	return r
}
