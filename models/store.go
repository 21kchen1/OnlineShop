package models

import (
	mysql "onlineshop/mysql"

	"github.com/jinzhu/gorm"
)

/**
 * @File : store.go
 * @Description : 订单模型 与其增删查改
 * @Author : chen
 * @Date : 2023/11/15
 */

type Store struct {
	gorm.Model
	UserID       int    `json:"userId"`
	StoreName    string `json:"storeName"`
	ContactInfo  string `json:"contactInfo"`
	Followers    int    `json:"followers"`
	ProductCount int    `json:"productCount"`
	SalesNum     int    `json:"saleNum"`
}

// 创建 Store
func CreateAStore(theStore *Store) (err error) {
	err = mysql.DB.Create(&theStore).Error

	return err
}

// 获取所有 Store
func GetAllStore() (theStoreList []*Store, err error) {
	err = mysql.DB.Find(&theStoreList).Error

	if err != nil {
		return nil, err
	}

	return theStoreList, nil
}

// 通过 id 获取 Store
func GetStoreByID(id int) (theStore Store, err error) {
	err = mysql.DB.Where("id = ?", id).First(&theStore).Error

	if err != nil {
		return theStore, err
	}

	return theStore, nil
}

// 更新一个存在于数据库的 Store
func UpdateAStore(theStore *Store) (err error) {
	err = mysql.DB.Save(&theStore).Error

	return err
}

// 通过 id 删除 Store
func DeleteStoreByID(id int) (err error) {
	err = mysql.DB.Where("id = ?", id).Delete(Store{}).Error

	return err
}

// GetStoreByUserID 根据用户ID获取商铺信息
func GetStoreByUserID(userID int) (store Store, err error) {
	err = mysql.DB.Where("user_id = ?", userID).First(&store).Error
	return store, err
}
