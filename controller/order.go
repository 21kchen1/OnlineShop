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

// EditOrder 修改订单信息接口
func EditOrder(c *gin.Context) {
	var requestData struct {
		OrderID      int    `json:"orderId"`
		OrderNumber  int    `json:"orderNumber"`
		OrderAddress string `json:"orderAddress"`
		OrderStatus  int    `json:"deliveryStatus"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	// 调用 service 修改订单信息
	err := service.EditOrder(requestData.OrderID, requestData.OrderNumber, requestData.OrderStatus)
	if err != nil {
		util.ErrRespon(c, err, "修改订单信息失败")
		return
	}

	// 返回成功信息
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "修改订单信息成功",
	})
}

// 第14接口，已完成获取订单model，商品信息未获取
// GetUserOrderList 获取用户订单列表
func GetUserOrderList(c *gin.Context) {
	var requestData struct {
		UserId int `json:"userId"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	// 调用 service 层函数获取用户订单列表
	orders, err := service.GetUserOrders(requestData.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "服务器错误",
		})
		return
	}

	// 构造返回参数
	var responseData []gin.H
	for _, order := range orders {
		orderData := gin.H{
			"productId":      order.ProductID,
			"productName":    order.ProductID,
			"orderNumber":    order.ProductID,
			"price":          order.ProductID,
			"deliveryStatus": order.OrderStatus,
			"date":           order.ProductID,
		}
		responseData = append(responseData, orderData)
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"data":      responseData,
	})
}
