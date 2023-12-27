package test

/**
 * @File : testProduct.go
 * @Description : 产品相关功能测试
 * @Author : chen
 * @Date : 2023/12/27
 */

import (
	"onlineshop/models"
	"onlineshop/service"
)

var testProductList = []models.Product{
	{
		StoreId:       123,
		ProductName:   "鞋子",
		Description:   "白",
		ProductStatus: 0,
		MonthNum:      100,
		Stock:         100,
		ProductType:   0,
		Likes:         1000,
		Comments:      1100,
		Price:         1,
	},
	{
		StoreId:       123,
		ProductName:   "鞋子",
		Description:   "黑",
		ProductStatus: 0,
		MonthNum:      110,
		Stock:         120,
		ProductType:   0,
		Likes:         1300,
		Comments:      1400,
		Price:         2,
	},
	{
		StoreId:       123,
		ProductName:   "衣服",
		Description:   "白",
		ProductStatus: 0,
		MonthNum:      100,
		Stock:         130,
		ProductType:   2,
		Likes:         14500,
		Comments:      1150,
		Price:         6,
	},
}

func AddProductData() {
	for i := range testProductList {
		service.AddProduct(&testProductList[i])
	}
}
