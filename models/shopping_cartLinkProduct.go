package models

import (
	mysql "onlineshop/mysql"

	"github.com/jinzhu/gorm"
)

/**
 * @File : shopping_cartLinkProduct.go
 * @Description : 购物车-商品
 * @Author : lei
 * @Date : 2023/12/31
 */

type ShoppingCartLinkProduct struct {
	gorm.Model
	ShoppingCartID int `json:"shoppingCartId" gorm:"foreignKey:ShoppingCartID"`
	ProductID      int `json:"productId" gorm:"foreignKey:ProductID"`
	Quantity       int `json:"quantity"`
}

// 购物车添加商品
func AddProductToShoppingCart(theShoppingCartLinkProduct *ShoppingCartLinkProduct) (err error) {
	err = mysql.DB.Create(&theShoppingCartLinkProduct).Error
	return err
}

// 根据购物车id搜索所有商品
func GetShoppingCartLinkProductByShoppingCartID(shoppingCartID int) (itemList []*ShoppingCartLinkProduct, err error) {
	err = mysql.DB.Where("shopping_cart_id = ?", shoppingCartID).Find(&itemList).Error
	return itemList, err
}

// 根据购物车id和商品id删除商品
func DeleteShoppingCartLinkProductByCartIDAndProductID(shoppingCartID int, productID int) (err error) {
	// 查询
	query := mysql.DB

	query = query.Where("shopping_cart_id = ?", shoppingCartID)
	query = query.Where("product_id = ?", productID)

	// 删除
	err = query.Delete(ShoppingCartLinkProduct{}).Error

	return err
}
