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
		FavoritesName: "114514",
		Count: 15,
	},
	{
		UserID: 1,
		FavoritesName: "11454",
		Count: 20,
	},
	{
		UserID: 1,
		FavoritesName: "11514",
		Count: 30,
	},
}

func AddFavoData() {
	for i := range testFavoList {
		service.AddFavorites(testFavoList[i])
	}
}