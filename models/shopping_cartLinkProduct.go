package models

import (
	"errors"
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

// CheckShoppingCartProductLinkExists 检查 shopping_cart_id - product_id 是否存在
func CheckShoppingCartProductLinkExists(shoppingCartID, productID int) (bool, error) {
	var link ShoppingCartLinkProduct

	err := mysql.DB.Where("shopping_cart_id = ? AND product_id = ?", shoppingCartID, productID).First(&link).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// UpdateShoppingCartProductQuantity 更新购物车商品项数量
func UpdateShoppingCartProductQuantity(shoppingCartID, productID, quantity int) error {
	var link ShoppingCartLinkProduct

	err := mysql.DB.Where("shopping_cart_id = ? AND product_id = ?", shoppingCartID, productID).First(&link).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("购物车项不存在")
		}
		return err
	}

	link.Quantity = quantity
	return mysql.DB.Save(&link).Error
}
