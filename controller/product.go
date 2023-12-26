/**
 * @File : product.go
 * @Description : 商品管理的控制器
 * @Author : you
 * @Date : 2023-12-18
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlineshop/service"
	"strconv"
)

// GetProductList 获取商品列表接口
func GetProductList(c *gin.Context) {
	var requestData struct {
		SearchKey   string `json:"searchKey"`
		ProductType int    `json:"type"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "获取参数失败",
		})
		return
	}

	// 调用服务函数获取商品列表
	productList, err := service.GetProductList(requestData.SearchKey, strconv.Itoa(requestData.ProductType))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "获取商品列表失败",
		})
		return
	}

	// 返回商品列表
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "获取商品列表成功",
		"data":      productList,
	})
}

/**
 * @File : product.go
 * @Description : 获取商品信息
 * @Author : chen
 * @Date : 2023-12-26
 */
func GetProduct(c *gin.Context) {
	var requestData struct {
		ProductId int `json:"productId"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "获取参数失败",
		})
		return
	}

	// 获取商品
	product, err := service.GetProduct(requestData.ProductId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "获取商品失败",
		})
		return
	}

	// 返回商品列表
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "获取商品成功",
		"data":      product,
	})
}

/**
 * @File : product.go
 * @Description : 根据id删除商品
 * @Author : chen
 * @Date : 2023-12-26
 */
func DeleteProduct(c *gin.Context) {
	var requestData struct {
		ProductId int `json:"productId"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "获取参数失败",
		})
		return
	}

	err := service.DeleteProduct(requestData.ProductId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "删除商品失败",
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"isSuccess": true,
		"msg":       "删除商品成功",
	})
}

/**
 * @File : product.go
 * @Description : 根据id获得商品数量
 * @Author : chen
 * @Date : 2023-12-26
 */
func GetProductNum(c *gin.Context) {
	var requestData struct {
		ProductId int `json:"productId"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "获取参数失败",
		})
		return
	}

	stock, err := service.GetProductNum(requestData.ProductId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "获取商品数量失败",
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"isSuccess": true,
		"msg":       "获取商品数量成功",
		"quantity":  stock,
	})
}
