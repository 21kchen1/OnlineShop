package models

import (
	"github.com/jinzhu/gorm"
	mysql "onlineshop/mysql"
)

/**
 * @File : favorites.go
 * @Description : 收藏模型 与其增删查改
 * @Author : chen
 * @Date : 2023/11/18
 */

type Favorites struct {
	gorm.Model
	UserID        int    `json:"userId"`        // 用户 ID
	FavoritesName string `json:"favoritesName"` // 收藏夹名称
	Count         int    `json:"count"`         // 收藏数量
}

// 创建 Favorites
func CreateAFavorites(theFavorites *Favorites) (err error) {
	err = mysql.DB.Create(&theFavorites).Error

	return err
}

// 获取所有 Favorites
func GetAllFavorites() (theFavoritesList []*Favorites, err error) {
	err = mysql.DB.Find(&theFavoritesList).Error

	if err != nil {
		return nil, err
	}

	return theFavoritesList, nil
}

// 通过 id 获取 Favorites
func GetFavoritesByID(id int) (theFavorites *Favorites, err error) {
	err = mysql.DB.Where("id = ?", id).First(&theFavorites).Error

	if err != nil {
		return nil, err
	}

	return theFavorites, nil
}

// 更新一个存在于数据库的 Favorites
func UpdateAFavorites(theFavorites *Favorites) (err error) {
	err = mysql.DB.Save(&theFavorites).Error

	return err
}

// 通过 id 删除 Favorites
func DeleteFavoritesByID(id int) (err error) {
	err = mysql.DB.Where("id = ?", id).Delete(Favorites{}).Error

	return err
}

// 通过用户id 获取对应 Favorites 列表
func GetFavoListByUserId(userId int) (theFavoritesList []*Favorites, err error) {
	err = mysql.DB.Where("user_id = ?", userId).Find(&theFavoritesList).Error

	return
}
