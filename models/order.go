package models

import (
	mysql "onlineshop/mysql"

	"github.com/jinzhu/gorm"
)

/**
 * @File : order.go
 * @Description : 订单模型 与其增删查改
 * @Author : chen
 * @Date : 2023/11/14
 */

type Order struct {
	gorm.Model
	OrderID        int    `json:"orderId"`
	UserID         int    `json:"userId"`
	ProductID      int    `json:"productId"`
	OrderStatus    int    `json:"orderStatus"`
	TotalPrice     int    `json:"totalPrice"`
	Quantity       int    `json:"quantity"`
	OrderTime      string `json:"orderTime"`
	PayTime        string `json:"payTime"`
	ShippingTime   string `json:"shippingTime"`
	CompletionTime string `json:"completionTime"`
}

// 创建 Order
func CreateAOrder(theOrder *Order) (err error) {
	err = mysql.DB.Create(&theOrder).Error

	return err
}

// 获取所有 Order
func GetAllOrder() (theOrderList []*Order, err error) {
	err = mysql.DB.Find(&theOrderList).Error

	if err != nil {
		return nil, err
	}

	return theOrderList, nil
}

// 通过 id 获取 Order
func GetOrderByID(id int) (theOrder Order, err error) {
	err = mysql.DB.Where("id = ?", id).First(&theOrder).Error

	if err != nil {
		return theOrder, err
	}

	return theOrder, nil
}

// 更新一个存在于数据库的 Order
func UpdateAOrder(theOrder *Order) (err error) {
	err = mysql.DB.Save(&theOrder).Error

	return err
}

// 通过 id 删除 Order
func DeleteOrderByID(id int) (err error) {
	err = mysql.DB.Where("id = ?", id).Delete(Order{}).Error

	return err
}

// GetOrderList 获取订单列表服务函数
func GetOrderList() (orders []*Order, err error) {
	// 查询数据库获取订单列表
	err = mysql.DB.Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// GetUserOrders 获取用户订单列表
func GetUserOrders(userID int) (orders []*Order, err error) {
	err = mysql.DB.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}
