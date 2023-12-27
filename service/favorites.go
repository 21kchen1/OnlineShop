package service

import (
	"errors"
	"onlineshop/models"
)

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
	for i := range favList {
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
	if theFavorites.FavoritesName == "" {
		err = errors.New("收藏夹名称不可为空")
		return
	}
	err = models.CreateAFavorites(&theFavorites)

	return
}

/**
 * @File : favorites.go
 * @Description : 删除收藏夹
 * @Author : chen
 * @Date : 2023/12/27
 */
func DeleteFavorites(favoId int) (err error) {
	err = models.DeleteFavoritesByID(favoId)

	return
}

/**
 * @File : favorites.go
 * @Description : 更新收藏夹名称
 * @Author : chen
 * @Date : 2023/12/27
 */
func UpdataFavoName(favoritesId int, FavoritesName string) (err error) {
	if FavoritesName == "" {
		err = errors.New("收藏夹名称不可为空")
		return
	}
	// 获取对应收藏夹
	theFavorites, err := models.GetFavoritesByID(favoritesId)

	if err != nil {
		return
	}

	theFavorites.FavoritesName = FavoritesName
	// 更新收藏夹
	err = models.UpdateAFavorites(&theFavorites)

	return
}