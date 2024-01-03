package test

import (
	"onlineshop/models"
)

var test = models.UserLinkShoppingcart {
	UserID: 1,
	ShoppingCartID: 1,
}

var ttest = []models.ShoppingCartLinkProduct {
	{
		ShoppingCartID: 1,
		ProductID: 1,
		Quantity: 1,
	},
	{
		ShoppingCartID: 1,
		ProductID: 2,
		Quantity: 2,
	},
	{
		ShoppingCartID: 1,
		ProductID: 8,
		Quantity: 3,
	},
}

func addShopping()  {
	models.CreateUserLinkShoppingcart(&test)
	for _, i := range ttest {
		models.AddProductToShoppingCart(&i)
	}
}