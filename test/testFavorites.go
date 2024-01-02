package test

import (
	"onlineshop/models"
	"onlineshop/service"
)

/**
 * @File : testFavorites.go
 * @Description : 收藏夹相关功能测试
 * @Author : chen
 * @Date : 2023/12/27
 */

var testFavoList = []models.Favorites {
	{
		UserID: 1,
		FavoritesName: "衣服",
		Count: 4,
	},
	{
		UserID: 1,
		FavoritesName: "裤子",
		Count: 4,
	},
	{
		UserID: 1,
		FavoritesName: "帽子",
		Count: 4,
	},
}

var testFPList = []models.FavoritesLinkProduct {
	{
		FavoritesID: 1,
		ProductID: 5,
	},
	{
		FavoritesID: 1,
		ProductID: 6,
	},
	{
		FavoritesID: 1,
		ProductID: 7,
	},
	{
		FavoritesID: 1,
		ProductID: 8,
	},
	{
		FavoritesID: 2,
		ProductID: 9,
	},
	{
		FavoritesID: 2,
		ProductID: 10,
	},
	{
		FavoritesID: 2,
		ProductID: 11,
	},
	{
		FavoritesID: 2,
		ProductID: 12,
	},
	{
		FavoritesID: 3,
		ProductID: 13,
	},
	{
		FavoritesID: 3,
		ProductID: 14,
	},
	{
		FavoritesID: 3,
		ProductID: 15,
	},
	{
		FavoritesID: 3,
		ProductID: 16,
	},
}

func addFavoData() {
	for i := range testFavoList {
		service.AddFavorites(testFavoList[i])
	}

	for i := range testFPList {
		service.AddFavoProduct(testFPList[i])
	}
}