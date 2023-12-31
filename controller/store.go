package controller

import (
	"net/http"
	"onlineshop/service"
	"onlineshop/util"

	"github.com/gin-gonic/gin"
)

// GetStoreID 获取商铺ID接口
func GetStoreID(c *gin.Context) {
	var requestData struct {
		UserID int `json:"userId"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	// service获取商铺ID
	storeID, err := service.GetStoreIDByUserID(requestData.UserID)
	if err != nil {
		util.ErrRespon(c, err, "获取商铺ID失败")
		return
	}

	// 返回商铺ID
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "获取商铺ID成功",
		"data":      map[string]interface{}{"shopid": storeID},
	})
}

// GetProductsByStoreID 获取商铺下的商品接口
func GetProductsByStoreID(c *gin.Context) {
	var requestData struct {
		StoreID int `json:"shopId"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	// service获取商铺下的商品
	products, err := service.GetProductsByStoreID(requestData.StoreID)
	if err != nil {
		util.ErrRespon(c, err, "获取商铺下的商品失败")
		return
	}

	// 返回商铺下的商品列表
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "获取商铺下的商品成功",
		"data":      products,
	})
}
