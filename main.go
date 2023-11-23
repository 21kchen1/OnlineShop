package main

import (
	"fmt"
	"onlineshop/models"
	"onlineshop/mysql"
	"github.com/gin-gonic/gin"
)

func main()  {
	// 数据库链接
	err := mysql.SqlInit()
	if err != nil {
		print("SqlInit Err ",err)
	}
	// 自动迁移
	mysql.DB.AutoMigrate(&models.Comment{}, &models.Favorites{}, &models.Log{}, &models.Order{}, &models.Product{}, &models.Store{}, &models.User{})
	// 执行结束关闭数据库
	defer mysql.DB.Close()
	r := gin.Default()


	// 路由组 具体待定
	userGroup := r.Group("/user")
	{
		userGroup.GET("/login", func(ctx *gin.Context) {})
	}

	//连接接口
	//打开cmd,输入ipconfig,找到ipv4地址
	err = r.Run("192.168.1.103:8080")
	if err != nil {
		fmt.Println("Gin Err ", err)
	}
}