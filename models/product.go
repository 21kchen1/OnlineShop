package models

import (
	"github.com/jinzhu/gorm"
	mysql "onlineshop/mysql"
)

/**
 * @File : product.go
 * @Description : 商品模型 与其增删查改
 * @Author : chen
 * @Date : 2023/11/13
 */

type Product struct {
	gorm.Model
	StoreId       uint   `json:"storeId"`
	ProductName   string `json:"productName"`
	Description   string `json:"description"`
	ProductStatus int    `json:"productStatus"`
	MonthNum      int64  `json:"monthNum"`
	Stock         int64  `json:"stock"`
	ProductType   int    `json:"productType"`
	Likes         int64  `json:"likes"`
	Comments      int64  `json:"comments"`
	Price         int64  `json:"price"`
}

// 创建 Product
func CreateAProduct(theProduct *Product) (err error) {
	err = mysql.DB.Create(&theProduct).Error

	return err
}

// 获取所有 Product
func GetAllProduct() (theProductList []*Product, err error) {
	err = mysql.DB.Find(&theProductList).Error

	if err != nil {
		return nil, err
	}

	return theProductList, nil
}

// 通过 id 获取 Product
func GetProductByID(id int) (theProduct *Product, err error) {
	err = mysql.DB.Where("id = ?", id).First(&theProduct).Error

	if err != nil {
		return nil, err
	}

	return theProduct, nil
}

// 更新一个存在于数据库的 Product
func UpdateAProduct(theProduct *Product) (err error) {
	err = mysql.DB.Save(&theProduct).Error

	return err
}

// 通过 id 删除 Product
func DeleteProductByID(id int) (err error) {
	err = mysql.DB.Where("id = ?", id).Delete(Product{}).Error

	return err
}
