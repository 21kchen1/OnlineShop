package main

import (
	"fmt"
	"onlineshop/models"
	"onlineshop/mysql"
	"onlineshop/routers"
	"onlineshop/test"
)

func main() {
	// 数据库链接
	err := mysql.SqlInit()
	if err != nil {
		print("SqlInit Err ", err)
	}
	// 自动迁移
	mysql.DB.AutoMigrate(&models.Comment{}, &models.Favorites{}, &models.Log{}, &models.Order{}, &models.Product{}, &models.Store{}, &models.User{}, &models.FavoritesLinkProduct{}, &models.Seller{})
	// 购物车与连接，user-购物车补充 from乐
	mysql.DB.AutoMigrate(&models.ShoppingCart{}, &models.ShoppingCartLinkProduct{}, &models.UserLinkShoppingcart{})
	// 执行结束关闭数据库
	defer mysql.DB.Close()

	// 输入测试数据
	test.TestAll()

	r := routers.SetupRouters()
	//连接接口
	//打开cmd,输入ipconfig,找到ipv4地址
	err = r.Run("localhost:5050")
	if err != nil {
		fmt.Println("Gin Err ", err)
	}
}
