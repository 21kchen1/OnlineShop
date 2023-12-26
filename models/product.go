package models

import (
	"fmt"
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
	StoreId       uint   `json:"storeId"`       // 店铺id
	ProductName   string `json:"productName"`   // 商品名称
	Description   string `json:"description"`   // 商品描述
	ProductStatus int    `json:"productStatus"` // 商品状态
	MonthNum      int64  `json:"monthNum"`      // 月销量
	Stock         int64  `json:"stock"`         // 库存数量
	ProductType   int    `json:"productType"`   // 商品类型
	Likes         int64  `json:"likes"`         // 商品喜欢数量
	Comments      int64  `json:"comments"`      // 商品评论数量
	Price         int64  `json:"price"`         // 商品价格
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

// GetProductList 获取商品列表
func GetProductList(searchKey, productType string) (products []*Product, err error) {
	// 构造查询条件
	query := mysql.DB

	fmt.Printf("Search Key: %s, Product Type: %s\n", searchKey, productType)

	// 判断是否有搜索关键词，有则添加条件
	if searchKey != "" {
		query = query.Where("product_name LIKE ?", "%"+searchKey+"%")
	}

	// 判断是否有商品类型，有则添加条件
	if productType != "" {
		query = query.Where("product_type = ?", productType)
	}

	// 查询数据库获取商品列表
	err = query.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
