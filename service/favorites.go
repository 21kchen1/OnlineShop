package service

import (
	"errors"
	"onlineshop/models"

	"github.com/jinzhu/gorm"
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
func GetFavoritesList(userId int) (refavList []map[string]interface{}, err error) {
	favList, err := models.GetFavoListByUserId(userId)

	if err != nil {
		return
	}

	// 构造
	for _, fav := range favList {
		item := map[string]interface{} {
			"id": fav.ID,
			"name": fav.FavoritesName,
			"num": fav.Count,
		}
		refavList = append(refavList, item)
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

/**
 * @File : favorites.go
 * @Description : 更新收藏夹名称
 * @Author : chen
 * @Date : 2023/12/27
 */
func GetFavoProductList(favoritesId int) (favProductIdList []map[string]interface{}, err error) {
	itemList, err := models.GetFavoritesLinkProductByFavoritesId(favoritesId)

	if err != nil {
		return nil, err
	}

	// 生成 id 列表
	for i := range itemList {
		productId := map[string]interface{} {
			"productId": itemList[i].ProductID,
		}
		favProductIdList = append(favProductIdList, productId)
	}

	return
}

/**
 * @File : favorites.go
 * @Description : 添加收藏夹商品
 * @Author : chen
 * @Date : 2023/12/27
 */
func AddFavoProduct(favoLinkProduct models.FavoritesLinkProduct) (err error) {
	if favoLinkProduct.FavoritesID == 0 || favoLinkProduct.ProductID == 0 {
		err = errors.New("收藏夹物品信息缺失")
		return
	}

	// 检测是否存在收藏夹
	theFavorites, err := models.GetFavoritesByID(favoLinkProduct.FavoritesID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	// 检测是否存在商品
	_, err = models.GetProductByID(favoLinkProduct.ProductID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	// 检测是否重复
	favProductIdList, err := GetFavoProductList(favoLinkProduct.FavoritesID)
	if err != nil {
		return
	}

	for i := range favProductIdList {
		if favProductIdList[i]["productId"] != favoLinkProduct.ProductID {
			continue
		}
		err = errors.New("收藏夹物品重复")
		return
	}

	err = models.AddProductToFavorites(&favoLinkProduct)
	if err != nil {
		return
	}
	// 收藏夹收藏数量更新
	theFavorites.Count = theFavorites.Count + 1
	// 保存收藏夹
	err = models.UpdateAFavorites(&theFavorites)

	return
}

/**
 * @File : favorites.go
 * @Description : 删除收藏夹商品
 * @Author : chen
 * @Date : 2023/12/27
 */
func DeleteFavoProduct(favoLinkProduct models.FavoritesLinkProduct) (err error) {
	if favoLinkProduct.FavoritesID == 0 || favoLinkProduct.ProductID == 0 {
		err = errors.New("收藏夹物品信息缺失")
		return
	}

	// 检测是否存在收藏夹
	theFavorites, err := models.GetFavoritesByID(favoLinkProduct.FavoritesID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	// 检测收藏夹是否存在商品
	favProductIdList, err := GetFavoProductList(favoLinkProduct.FavoritesID)
	if err != nil {
		return
	}

	for i := range favProductIdList {
		if favProductIdList[i]["productId"] == favoLinkProduct.ProductID {
			break
		}
		if i != len(favProductIdList) - 1 {
			continue
		}
		// 到最后一个也没有找到
		err = errors.New("收藏夹不存在物品")
		return
	}

	err = models.DeleteFavoritesLinkProductByFIdAndPId(favoLinkProduct.FavoritesID, favoLinkProduct.ProductID)
	if err != nil {
		return
	}
	// 收藏夹收藏数量更新
	theFavorites.Count = theFavorites.Count - 1
	// 保存收藏夹
	err = models.UpdateAFavorites(&theFavorites)

	return err
}