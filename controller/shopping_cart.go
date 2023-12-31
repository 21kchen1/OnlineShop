package controller

/**
 * @File : shopping_cart.go
 * @Description : 购物车管理的控制器
 * @Author : lei
 * @Date : 2023-12-31
 */

import (
	"net/http"
	"onlineshop/service"
	"onlineshop/util"

	"github.com/gin-gonic/gin"
)

// GetShopIDByShopName 根据商铺名称获取商铺ID
func GetShopIDByShopName(c *gin.Context) {
	// 从请求中获取商铺名称
	var requestData struct {
		ShopName string `json:"shopName"`
	}
	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}
	// 调用 service 层函数获取商铺ID
	shopID, err := service.GetShopIDByShopName(requestData.ShopName)
	if err != nil {
		util.ErrRespon(c, err, "获取商铺ID失败")
		return
	}
	// 成功获取商铺ID，返回成功信息和商铺ID
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "成功获取商铺ID",
		"shopId":    shopID,
	})
}

// GetProductIDByProductName 根据商品名称获取商品ID
func GetProductIDByProductName(c *gin.Context) {
	// 从请求中获取商铺名称
	var requestData struct {
		ProductName string `json:"productName"`
	}
	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}
	// 调用 service 层函数获取商铺ID
	shopID, err := service.GetProductIDByProductName(requestData.ProductName)
	if err != nil {
		util.ErrRespon(c, err, "获取商品ID失败")
		return
	}
	// 成功获取商铺ID，返回成功信息和商铺ID
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "成功获取商品ID",
		"shopId":    shopID,
	})
}
