package models

import (
	"github.com/jinzhu/gorm"
	mysql "onlineshop/mysql"
)

/**
 * @File : favoritesLinkProduct.go
 * @Description : 收藏-商品
 * @Author : chen
 * @Date : 2023/12/26
 */

type FavoritesLinkProduct struct {
	gorm.Model
	FavoritesID string `json:"favoritesId"`
	ProductID   string `json:"productId"`
}

// 收藏夹添加商品
func AddProductToFavorites(theFavoritesLinkProduct *FavoritesLinkProduct) (err error) {
	err = mysql.DB.Create(&theFavoritesLinkProduct).Error

	return
}

// 根据收藏夹id搜索所有
func GetFavoritesLinkProductByFavoritesId(favoritesId int) (itemList []*FavoritesLinkProduct, err error) {
	err = mysql.DB.Where("favoritesid = ", favoritesId).Find(&itemList).Error

	return
}

// 根据收藏夹id删除商品
