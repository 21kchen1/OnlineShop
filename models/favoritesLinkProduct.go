package models

import (
	mysql "onlineshop/mysql"
	"github.com/jinzhu/gorm"
)

/**
 * @File : favoritesLinkProduct.go
 * @Description : 收藏-商品
 * @Author : chen
 * @Date : 2023/12/26
 */

type FavoritesLinkProduct struct {
	gorm.Model
	FavoritesID int `json:"favoritesId" gorm:"foreignKey:FavoritesID"`
	ProductID   int `json:"productId" gorm:"foreignKey:ProductID"`
}

// 收藏夹添加商品
func AddProductToFavorites(theFavoritesLinkProduct *FavoritesLinkProduct) (err error) {
	err = mysql.DB.Create(&theFavoritesLinkProduct).Error

	return
}

// 根据收藏夹id搜索所有
func GetFavoritesLinkProductByFavoritesId(favoritesId int) (itemList []*FavoritesLinkProduct, err error) {
	err = mysql.DB.Where("favorites_id = ?", favoritesId).Find(&itemList).Error

	return
}

// 根据收藏夹id删除商品
func DeleteFavoritesLinkProductByFIdAndPId(favoritesId int, productId int) (err error) {
	// 查询
	query := mysql.DB

	query = query.Where("favorites_id = ?", favoritesId)
	query = query.Where("product_id = ?", productId)

	// 删除
	err = query.Delete(FavoritesLinkProduct{}).Error

	return err
}
