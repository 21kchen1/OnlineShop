package test

import "onlineshop/models"

/**
 * @File : testOrder.go
 * @Description : 订单相关功能测试
 * @Author : chen
 * @Date : 2023/12/29
 */

var testOrderList = []models.Order {
	{
		ProductID: 1,
		TotalPrice: 2,
	},
	{
		ProductID: 3,
		TotalPrice: 4,
	},
}

func AddOrder() {
	for _, itme := range testOrderList {
		models.CreateAOrder(&itme)
	}
}

func Block() {

}