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
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	PhoneNum string `json:"phone_num"`
	UserType int    `json:"user_type"`
	Email    string `json:"email"`
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
