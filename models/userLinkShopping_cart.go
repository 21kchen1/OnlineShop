package models

import (
	"errors"
	mysql "onlineshop/mysql"

	"github.com/jinzhu/gorm"
)

/**
 * @File : userLinkShopping_cart.go
 * @Description : 用户-购物车
 * @Author : lei
 * @Date : 2023/12/31
 */

// UserLinkShoppingcart 用户-购物车关联
type UserLinkShoppingcart struct {
	gorm.Model
	UserID         int `json:"userId" gorm:"foreignKey:UserID"`
	ShoppingCartID int `json:"shoppingCartId" gorm:"foreignKey:ShoppingCartID"`
}

// 创建用户购物车链接
func CreateUserLinkShoppingcart(theUserLinkShoppingcart *UserLinkShoppingcart) (err error) {
	err = mysql.DB.Create(&theUserLinkShoppingcart).Error
	return err
}

// 获取用户购物车链接信息
func GetUserLinkShoppingcartByUserID(userID int) (theUserLinkShoppingcart UserLinkShoppingcart, err error) {
	err = mysql.DB.Where("user_id = ?", userID).First(&theUserLinkShoppingcart).Error
	return theUserLinkShoppingcart, err
}

// 更新用户购物车链接
func UpdateUserLinkShoppingcart(theUserLinkShoppingcart *UserLinkShoppingcart) (err error) {
	err = mysql.DB.Save(&theUserLinkShoppingcart).Error
	return err
}

// 删除用户购物车链接
func DeleteUserLinkShoppingcart(userID int) (err error) {
	err = mysql.DB.Where("user_id = ?", userID).Delete(UserLinkShoppingcart{}).Error
	return err
}

// CheckUserShoppingCartLinkExists 检查 user_id - shopping_cart_id 是否存在
func CheckUserShoppingCartLinkExists(userID, shoppingCartID int) (bool, error) {
	var link UserLinkShoppingcart

	err := mysql.DB.Where("user_id = ? AND shopping_cart_id = ?", userID, shoppingCartID).First(&link).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
