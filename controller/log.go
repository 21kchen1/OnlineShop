package controller

/**
 * @File : log.go
 * @Description : 日志管理的控制器
 * @Author : lei
 * @Date : 2023-12-29
 */
import (
	"net/http"
	"onlineshop/models"
	"onlineshop/service"
	"onlineshop/util"

	"github.com/gin-gonic/gin"
)

// GetLogList 获取日志列表接口
func GetLogList(c *gin.Context) {
	// 调用 service 获取日志列表
	logList, err := models.GetAllLog()
	if err != nil {
		util.ErrRespon(c, err, "获取日志列表失败")
		return
	}

	// 返回日志列表
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "获取日志列表成功",
		"data":      logList,
	})
}

// AddLog 添加日志接口
func AddLog(c *gin.Context) {
	var requestData struct {
		UserID  int    `json:"userId"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	// 调用 service 添加日志
	err := service.AddLog(requestData.UserID, requestData.Title, requestData.Content)
	if err != nil {
		util.ErrRespon(c, err, "添加日志失败")
		return
	}

	// 返回成功信息
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "添加日志成功",
	})
}

// DeleteLog 删除日志接口
func DeleteLog(c *gin.Context) {
	var requestData struct {
		LogID int `json:"logId"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	// 调用 service 删除日志
	err := service.DeleteLog(requestData.LogID)
	if err != nil {
		util.ErrRespon(c, err, "删除日志失败")
		return
	}

	// 返回成功信息
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "删除日志成功",
	})
}

// EditLog 修改日志接口
func EditLog(c *gin.Context) {
	var requestData struct {
		LogID   int    `json:"logId"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	// 调用 service 修改日志
	err := service.EditLog(requestData.LogID, requestData.Title, requestData.Content)
	if err != nil {
		util.ErrRespon(c, err, "修改日志失败")
		return
	}

	// 返回成功信息
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "修改日志成功",
	})
}
