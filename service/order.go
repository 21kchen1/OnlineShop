package service

import (
	"fmt"
	"onlineshop/models"
)

/**
 * @File : order.go
 * @Description : 订单相关的服务
 * @Author : lei
 * @Date : 2023/12/29
 */

// GetOrderList 获取所有订单列表服务函数
func GetOrderList() (orderList []map[string]interface{}, err error) {
	// 调用数据库模型的方法获取订单列表
	orders, err := models.GetOrderList()
	if err != nil {
		return nil, fmt.Errorf("获取订单列表失败：%w", err)
	}

	// 构造返回数据
	for _, order := range orders {
		orderData := map[string]interface{}{
			"OrderID":        order.ID,
			"ProductID":      order.ProductID,
			"OrderStatus":    order.OrderStatus,
			"TotalPrice":     order.TotalPrice,
			"Quantity":       order.Quantity,
			"OrderTime":      order.OrderTime,
			"PayTime":        order.PayTime,
			"ShippingTime":   order.ShippingTime,
			"CompletionTime": order.CompletionTime,
			// 添加其他需要返回的订单信息字段
		}
		orderList = append(orderList, orderData)
	}

	return orderList, nil
}

// EditOrder 修改订单信息服务函数
func EditOrder(orderID, orderNumber, deliveryStatus int) error {
	// 获取订单对象
	order, err := models.GetOrderByID(orderID)
	if err != nil {
		return fmt.Errorf("获取订单失败：%w", err)
	}

	// 更新订单信息
	order.OrderID = orderNumber
	order.OrderStatus = deliveryStatus

	// 调用数据库模型的方法保存更新后的订单信息
	err = models.UpdateAOrder(&order)
	if err != nil {
		return fmt.Errorf("修改订单信息失败：%w", err)
	}

	return nil
}

// GetUserOrders 获取用户订单列表
func GetUserOrders(userID int) ([]*models.Order, error) {
	if userID == 0 {
		return models.GetOrderList()
	}
	return models.GetUserOrders(userID)
}
