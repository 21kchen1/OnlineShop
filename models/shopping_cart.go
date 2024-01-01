package models

/**
 * @File : shopping_cart.go
 * @Description : 购物车模型 与其增删查改
 * @Author : lei
 * @Date : 2023/12/31
 */

import (
	"errors"
	mysql "onlineshop/mysql"

	"github.com/jinzhu/gorm"
)

type ShoppingCart struct {
	gorm.Model
	ShoppingCartName string `json:"ShoppingCartName"` // 购物车名称

}

// 创建 ShoppingCart
func CreateShoppingCart(shoppingCart *ShoppingCart) (err error) {
	err = mysql.DB.Create(shoppingCart).Error
	return err
}

// 获取所有 ShoppingCart
func GetAllShoppingCarts() (shoppingCarts []*ShoppingCart, err error) {
	err = mysql.DB.Find(&shoppingCarts).Error
	if err != nil {
		return nil, err
	}
	return shoppingCarts, nil
}

// 通过 id 获取 ShoppingCart
func GetShoppingCartByID(id int) (shoppingCart ShoppingCart, err error) {
	err = mysql.DB.Where("id = ?", id).First(&shoppingCart).Error
	if err != nil {
		return shoppingCart, err
	}
	return shoppingCart, nil
}

// 更新一个存在于数据库的 ShoppingCart
func UpdateShoppingCart(shoppingCart *ShoppingCart) (err error) {
	err = mysql.DB.Save(shoppingCart).Error
	return err
}

// 通过 id 删除 ShoppingCart
func DeleteShoppingCartByID(id int) (err error) {
	err = mysql.DB.Where("id = ?", id).Delete(ShoppingCart{}).Error
	return err
}

// GetShoppingCartLinkProductByUserIDShopIDProductID 根据用户ID、商铺ID和商品ID获取购物车商品项
func GetShoppingCartLinkProductByUserIDShopIDProductID(userID, shopID, productID int) (*ShoppingCartLinkProduct, error) {
	var cartItem ShoppingCartLinkProduct

	err := mysql.DB.Where("user_id = ? AND shop_id = ? AND product_id = ?", userID, shopID, productID).First(&cartItem).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("购物车项不存在")
		}
		return nil, err
	}

	return &cartItem, nil
}

// UpdateShoppingCartLinkProduct 更新购物车商品项
func UpdateShoppingCartLinkProduct(cartItem *ShoppingCartLinkProduct) error {
	return mysql.DB.Save(cartItem).Error
}
