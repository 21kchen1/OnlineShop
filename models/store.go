package models

import (
	"github.com/jinzhu/gorm"
	mysql "onlineshop/mysql"
)

/**
 * @File : store.go
 * @Description : 订单模型 与其增删查改
 * @Author : chen
 * @Date : 2023/11/15
 */

type Store struct {
	gorm.Model
	UserID       uint   `json:"userId"`
	StoreName    string `json:"storeName"`
	ContactInfo  string `json:"contactInfo"`
	Followers    int64  `json:"followers"`
	ProductCount int64  `json:"productCount"`
	SalesNum     int64  `json:"saleNum"`
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
func GetStoreByID(id int) (theStore *Store, err error) {
	err = mysql.DB.Where("id = ?", id).First(&theStore).Error

	if err != nil {
		return nil, err
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
