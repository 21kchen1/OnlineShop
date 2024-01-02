/**
 * @File : product.go
 * @Description : 商品管理的控制器
 * @Author : you
 * @Date : 2023-12-18
 */

package controller

import (
	"net/http"
	"onlineshop/models"
	"onlineshop/service"
	"onlineshop/util"
	"github.com/gin-gonic/gin"
)

// GetProductList 获取商品列表接口
func GetProductList(c *gin.Context) {
	var requestData struct {
		SearchKey   string `json:"searchKey"`
		ProductType int    `json:"type"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	// 调用服务函数获取商品列表
	productList, err := service.GetProductList(requestData.SearchKey, requestData.ProductType)
	if err != nil {
		util.ErrRespon(c, err, "获取商品列表失败")
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
 * @Description : 根据信息添加商品
 * @Author : chen
 * @Date : 2023-12-26
 */
func AddProduct(c *gin.Context) {
	var requestData models.Product

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	err := service.AddProduct(&requestData)

	if err != nil {
		util.ErrRespon(c, err, "添加商品失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "添加商品成功",
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
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	// 获取商品
	product, err := service.GetProduct(requestData.ProductId)

	if err != nil {
		util.ErrRespon(c, err, "获取商品失败")
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
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	err := service.DeleteProduct(requestData.ProductId)

	if err != nil {
		util.ErrRespon(c, err, "删除商品失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "删除商品成功",
	})
}

/**
 * @File : product.go
 * @Description : 根据id修改商品信息
 * @Author : chen
 * @Date : 2023-12-26
 */
func EditProduct(c *gin.Context) {
	var requestData struct {
		ProductId int `json:"productId"`
		models.Product
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	var theProduct models.Product
	{
	}
	theProduct = requestData.Product
	err := service.EditProduct(requestData.ProductId, theProduct)

	if err != nil {
		util.ErrRespon(c, err, "修改商品信息失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "修改商品信息成功",
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
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	stock, err := service.GetProductNum(requestData.ProductId)

	if err != nil {
		util.ErrRespon(c, err, "获取商品数量失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "获取商品数量成功",
		"quantity":  stock,
	})
}

/**
 * @File : product.go
 * @Description : 根据id修改商品数量
 * @Author : chen
 * @Date : 2023-12-26
 */
func EditProductNum(c *gin.Context) {
	var requestData struct {
		ProductId    int `json:"productId"`
		EditQuantity int `json:"editQuantity"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	err := service.EditProductNum(requestData.ProductId, requestData.EditQuantity)

	if err != nil {
		util.ErrRespon(c, err, "修改商品数量失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "修改商品数量成功",
	})
}

// GetRecommendedProducts 获取主页推荐商品
func GetRecommendedProducts(c *gin.Context) {
	// 调用 service 层获取推荐商品的函数
	recommendedProducts, err := service.GetRecommendedProducts()

	if err != nil {
		util.ErrRespon(c, err, "获取推荐商品失败")
		return
	}

	// 返回推荐商品列表
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "获取推荐商品成功",
		"data":      recommendedProducts,
	})
}
