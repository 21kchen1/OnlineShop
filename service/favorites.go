package service

import "onlineshop/models"

/**
 * @File : favorites.go
 * @Description : 收藏夹相关的服务
 * @Author : chen
 * @Date : 2023/12/27
 */

/**
 * @File : favorites.go
 * @Description : 根据用户id获取收藏夹列表
 * @Author : chen
 * @Date : 2023/12/27
 */
func GetFavoritesList(userId int) (favIdList []int, err error) {
	favList, err := models.GetFavoListByUserId(userId)

	if err != nil {
		return
	}

	// 构造
	for i := range favIdList {
		favIdList = append(favIdList, int(favList[i].ID))
	}

	return
}

/**
 * @File : favorites.go
 * @Description : 添加收藏夹
 * @Author : chen
 * @Date : 2023/12/27
 */

func AddFavorites(theFavorites models.Favorites) (err error) {
	err = models.CreateAFavorites(&theFavorites)

	return
}