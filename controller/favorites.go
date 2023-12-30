package controller

import (
	"net/http"
	"onlineshop/models"
	"onlineshop/service"
	"onlineshop/util"
	"github.com/gin-gonic/gin"
)

/**
 * @File : favorites.go
 * @Description : 收藏夹管理的控制器
 * @Author : chen
 * @Date : 2023-12-27
 */

/**
 * @File : favorites.go
 * @Description : 根据用户id获取收藏夹列表
 * @Author : chen
 * @Date : 2023-12-27
 */
func GetFavoritesList(c *gin.Context) {
	var requestData struct {
		UserId int `json:"userId"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	// 获取收藏夹列表
	favList, err := service.GetFavoritesList(requestData.UserId)

	if err != nil {
		util.ErrRespon(c, err, "获取收藏夹列表失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "获取收藏列表成功",
		"data":      favList,
	})
}

/**
 * @File : favorites.go
 * @Description : 创建收藏夹
 * @Author : chen
 * @Date : 2023-12-27
 */
func AddFavorites(c *gin.Context) {
	var requestData models.Favorites

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	// 创建收藏夹
	err := service.AddFavorites(requestData)

	if err != nil {
		util.ErrRespon(c, err, "添加收藏夹失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "添加收藏夹成功",
	})
}

/**
 * @File : favorites.go
 * @Description : 删除收藏夹
 * @Author : chen
 * @Date : 2023-12-27
 */
func DeleteFavorites(c *gin.Context) {
	var requestData struct {
		FavoritesId int `json:"favoritesId"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	// 删除收藏夹
	err := service.DeleteFavorites(requestData.FavoritesId)

	if err != nil {
		util.ErrRespon(c, err, "删除收藏夹失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "删除收藏夹成功",
	})
}

/**
 * @File : favorites.go
 * @Description : 更新收藏夹名称
 * @Author : chen
 * @Date : 2023-12-27
 */
func UpdataFavoName(c *gin.Context) {
	var requestData struct {
		FavoritesId   int    `json:"favoritesId"`
		FavoritesName string `json:"favoritesName"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	// 更新收藏夹名称
	err := service.UpdataFavoName(requestData.FavoritesId, requestData.FavoritesName)

	if err != nil {
		util.ErrRespon(c, err, "更新收藏夹名称失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "更新收藏夹名称成功",
	})
}

/**
 * @File : favorites.go
 * @Description : 获取收藏夹商品
 * @Author : chen
 * @Date : 2023-12-27
 */
func GetFavoProductList(c *gin.Context) {
	var requestData struct {
		FavoritesId int `json:"favoritesId"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	// 获取收藏夹商品列表
	favProductIdList, err := service.GetFavoProductList(requestData.FavoritesId)

	if err != nil {
		util.ErrRespon(c, err, "获取收藏夹商品id列表失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "获取收藏夹商品id列表成功",
		"data":      favProductIdList,
	})
}

/**
 * @File : favorites.go
 * @Description : 收藏夹添加商品
 * @Author : chen
 * @Date : 2023-12-27
 */
func AddFavoProduct(c *gin.Context) {
	var requestData models.FavoritesLinkProduct

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	// 添加收藏夹商品
	err := service.AddFavoProduct(requestData)

	if err != nil {
		util.ErrRespon(c, err, "添加收藏夹商品失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "添加收藏夹商品成功",
	})
}

/**
 * @File : favorites.go
 * @Description : 收藏夹删除商品
 * @Author : chen
 * @Date : 2023-12-27
 */
func DeleteFavoProduct(c *gin.Context) {
	var requestData models.FavoritesLinkProduct

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	// 删除收藏夹商品
	err := service.DeleteFavoProduct(requestData)

	if err != nil {
		util.ErrRespon(c, err, "删除收藏夹商品失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "删除收藏夹商品成功",
	})
}