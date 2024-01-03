package test

import (
	"onlineshop/models"
)

/**
 * @File : testStore.go
 * @Description : 商铺相关功能测试
 * @Author : chen
 * @Date : 2023/12/27
 */

var testStore = []models.Store{
	{
		UserID: 1,
		StoreName: "小二店",
		ContactInfo: "1234556",
		Followers: 1000,
		ProductCount: 123,
		SalesNum: 1000,
	},
	{
		UserID: 2,
		StoreName: "小四店",
		ContactInfo: "1234556",
		Followers: 1000,
		ProductCount: 123,
		SalesNum: 1000,
	},
}

func addStore() {
	for _, i := range testStore {
		models.CreateAStore(&i)
	}
}
