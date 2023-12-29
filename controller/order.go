package controller

/**
 * @File : order.go
 * @Description : 订单管理的控制器
 * @Author : lei
 * @Date : 2023-12-29
 */
import (
	"fmt"
	"net/http"
	"onlineshop/models"
	"onlineshop/service"
	"onlineshop/util"

	"github.com/gin-gonic/gin"
)

// GetOrderList 获取订单列表接口
func GetOrderList(c *gin.Context) {
	// service获取订单列表
	orderList, err := service.GetOrderList()
	if err != nil {
		util.ErrRespon(c, err, "获取订单列表失败")
		return
	}

	// 直接返回订单列表
	fmt.Println("Order List:", orderList)
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "获取订单列表成功",
		"data":      orderList,
	})

}

// DeleteOrder 删除订单接口
func DeleteOrder(c *gin.Context) {
	var requestData struct {
		OrderID int `json:"orderId"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	// 调用 models 删除订单
	err := models.DeleteOrderByID(requestData.OrderID)
	if err != nil {
		util.ErrRespon(c, err, "删除订单失败")
		return
	}

	// 返回成功信息
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "删除订单成功",
	})
}
