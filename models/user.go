package models

import (
	"github.com/jinzhu/gorm"
	mysql "onlineshop/mysql"
)

/**
 * @File : user.go
 * @Description : 用户模型 与其增删查改
 * @Author : chen
 * @Date : 2023/11/12
 */

type User struct {
	gorm.Model
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
	PhoneNum string `json:"phoneNum"`
	UserType int    `json:"userType"`
	Email    string `json:"email" binding:"required"`
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
func GetUserByID(id int) (theUser *User, err error) {
	err = mysql.DB.Where("id = ?", id).First(&theUser).Error

	if err != nil {
		return nil, err
	}

	return theUser, nil
}

// 通过 UserName 获取 user
func GetUserByName(userName string) (theUser *User, err error) {
	result := mysql.DB.Where("UserName = ?", userName).First(&theUser)
	err = result.Error

	if err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, nil
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
	err = mysql.DB.Where("id = ?", id).Delete(User{}).Error

	return err
}
