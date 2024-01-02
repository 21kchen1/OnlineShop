package test

import "onlineshop/models"

var testOrder = []models.Order {
	{
		OrderID: 1,
		UserID: 1,
		ProductID: 4,
		OrderStatus: 0,
		TotalPrice: 114514,
		Quantity: 1,
	},
	{
		OrderID: 2,
		UserID: 1,
		ProductID: 5,
		OrderStatus: 0,
		TotalPrice: 114515,
		Quantity: 1,
	},
	{
		OrderID: 2,
		UserID: 1,
		ProductID: 6,
		OrderStatus: 0,
		TotalPrice: 114516,
		Quantity: 1,
	},
}

func addOrder() {
	for _, i := range testOrder {
		models.CreateAOrder(&i)
	}
}