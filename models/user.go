package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	mysql "onlineshop/mysql"
)

/**
 * @File : user.go
 * @Description : 用户模型 与其增删查改
 * @Author : chen
 * @Date : 2023/11/12
 */

// Seller 商家信息模型

type Seller struct {
	gorm.Model
	UserID  uint   `json:"user_id"` // 关联的用户 ID
	Address string `json:"address"`
}

// 在 User 结构体中添加一个 Seller 字段，用于关联商家信息

type User struct {
	gorm.Model
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
	PhoneNum string `json:"phoneNum"`
	UserType int    `json:"userType"`
	Email    string `json:"email" binding:"required"`
	Seller   Seller // 关联的商家信息
}

// 创建 user
func CreateAUser(theUser *User) (err error) {
	err = mysql.DB.Create(&theUser).Error

	return err
}

// 获取所有 user
func GetAllUser() (theUserList []*User, err error) {
	err = mysql.DB.Find(&theUserList).Error

	if err != nil {
		return nil, err
	}

	return theUserList, nil
}

// 通过 id 获取 user
func GetUserByID(id int) (theUser User, err error) {
	err = mysql.DB.Where("id = ?", id).First(&theUser).Error

	if err != nil {
		return theUser, err
	}

	return theUser, nil
}

// 通过 UserName 获取 user
func GetUserByName(userName string) (theUser *User, err error) {
	theUser = new(User)
	err = mysql.DB.Where("user_name = ?", userName).First(&theUser).Error

	if err != nil {
		return nil, err
	}

	return theUser, nil
}

// 更新一个存在于数据库的 user
func UpdateAUser(theUser *User) (err error) {
	err = mysql.DB.Save(&theUser).Error

	return err
}

// 通过 id 删除 user
func DeleteUserByID(id int) (err error) {
	result := mysql.DB.Where("id = ?", id).Delete(User{})
	if result.Error != nil {
		fmt.Println("Error deleting user:", result.Error)
	}

	return result.Error
}

// 获取商家列表
func GetSellerList() (sellerList []*User, err error) {
	// 使用Preload方法来预加载关联的Seller信息
	err = mysql.DB.Where("user_type = ?", 1).Preload("Seller").Find(&sellerList).Error
	if err != nil {
		return nil, err
	}

	return sellerList, nil
}

// 级联删除商家地址信息
func DeleteSellerByID(id int) (err error) {
	// 首先找到关联的Seller信息
	var seller Seller
	if err := mysql.DB.Where("user_id = ?", id).First(&seller).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("Error finding seller:", err)
		return err
	}

	// 如果找到了Seller记录，则删除它
	if seller.ID != 0 {
		if err := mysql.DB.Delete(&seller).Error; err != nil {
			fmt.Println("Error deleting seller:", err)
			return err
		}
	}

	// 删除User记录
	if err := mysql.DB.Where("id = ?", id).Delete(User{}).Error; err != nil {
		fmt.Println("Error deleting user:", err)
		return err
	}

	return nil
}
