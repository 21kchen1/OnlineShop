package util

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

/**
 * @File : errRespon.go
 * @Description : 错误返回公共函数
 * @Author : chen
 * @Date : 2023/12/29
 */
func ErrRespon(c *gin.Context, err error, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"isSuccess": false,
		"msg":       message + " :" + err.Error(),
	})
}
