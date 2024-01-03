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

// GetShoppingCart 获取当前账号购物车数据
func GetShoppingCart(c *gin.Context) {
	// 从请求中获取用户ID
	var requestData struct {
		UserID int `json:"userId"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		// 处理参数绑定失败的情况
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "参数绑定失败",
			"data":      nil,
		})
		return
	}

	// 调用 service 层函数获取购物车数据
	cartData, err := service.GetShoppingCartDataByUserID(requestData.UserID)

	if err != nil {
		// 处理获取购物车数据失败的情况
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "获取购物车数据失败",
			"data":      nil,
		})
		return
	}

	// 处理成功，返回购物车数据
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "成功获取购物车数据",
		"data":      cartData,
	})
}

// CheckUserShoppingCartLinkExistsAndUpadteQuantity 检查 user_id - shopping_cart_id 是否存在，
// 检查 shopping_cart_id - product_id 是否存在并更新数量
func CheckUserShoppingCartLinkExistsAndUpadteQuantity(c *gin.Context) {
	var requestData struct {
		UserId    int `json:"userId"`
		ShopId    int `json:"shopId"`
		ProductId int `json:"productId"`
		Quantity  int `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	// 检查 user_id - shopping_cart_id 是否存在
	existsUserLink, err := service.CheckUserShoppingCartLinkExists(requestData.UserId, requestData.ShopId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       requestData.ShopId,
			"user":      requestData.UserId,
		})
		return
	}

	if !existsUserLink {
		c.JSON(http.StatusOK, gin.H{
			"isSuccess": false,
			"msg":       "用户购物车链接不存在",
		})
		return
	}

	// 检查 shopping_cart_id - product_id 是否存在并更新数量
	existsCartProductLink, err := service.CheckShoppingCartProductLinkExists(requestData.ShopId, requestData.ProductId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "服务器错误",
		})
		return
	}

	if !existsCartProductLink {
		c.JSON(http.StatusOK, gin.H{
			"isSuccess": false,
			"msg":       "购物车商品链接不存在",
		})
		return
	}

	// 更新购物车商品项数量
	err = service.UpdateShoppingCartProductQuantity(requestData.ShopId, requestData.ProductId, requestData.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "服务器错误",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "更新成功",
	})
}
