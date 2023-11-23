package models

import (
	"github.com/jinzhu/gorm"
	mysql "onlineshop/mysql"
)

/**
 * @File : order.go
 * @Description : 订单模型 与其增删查改
 * @Author : chen
 * @Date : 2023/11/14
 */

type Order struct {
	gorm.Model
	ProductID      uint   `json:"product_id"`
	OrderStatus    int    `json:"order_status"`
	TotalPrice     int    `json:"total_price"`
	Quantity       int    `json:"quantity"`
	OrderTime      string `json:"order_time"`
	PayTime        string `json:"pay_time"`
	ShippingTime   string `json:"shipping_time"`
	CompletionTime string `json:"completion_time"`
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
func GetOrderByID(id int) (theOrder *Order, err error) {
	err = mysql.DB.Where("id = ?", id).First(&theOrder).Error

	if err != nil {
		return nil, err
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
