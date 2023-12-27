package controller

import (
	"net/http"
	"onlineshop/models"
	"onlineshop/service"

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
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "获取参数失败",
		})
		return
	}

	// 获取收藏夹列表
	favIdList, err := service.GetFavoritesList(requestData.UserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "获取收藏列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "获取收藏列表成功",
		"data":      favIdList,
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
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "获取参数失败",
		})
		return
	}

	// 创建收藏夹
	err := service.AddFavorites(requestData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "添加收藏夹失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "添加收藏夹成功",
	})
}
