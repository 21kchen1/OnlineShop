package models

/**
 * @File : shopping_cart.go
 * @Description : 购物车模型 与其增删查改
 * @Author : lei
 * @Date : 2023/12/31
 */

import (
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
