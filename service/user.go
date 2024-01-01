package service

import (
	"errors"
	"onlineshop/models"
)

/**
 * @File : user.go
 * @Description : 用户相关的服务
 * @Author : chen
 * @Date : 2023/12/03
 */

func CreateUsers(theUser *models.User) (err error) {
	err = models.CreateAUser(theUser)

	return err
}

func CheckUserExistByName(theUser *models.User) (err error) {
	_, err = models.GetUserByName(theUser.UserName)

	if err == nil {
		return errors.New("Found")
	}

	if err.Error() == "record not found" {
		return nil
	}

	return err
}

// 用户登录验证
func CheckUserLogin(loginData *models.User) (err error) {
	// 检查用户名是否为空
	if loginData.UserName == "" {
		return errors.New("用户名不能为空")
	}

	// 检查密码是否为空
	if loginData.PassWord == "" {
		return errors.New("用户密码不能为空")
	}

	// 验证用户名是否存在
	user, err := models.GetUserByName(loginData.UserName)
	if err != nil {
		if err.Error() == "record not found" {
			return errors.New("用户名不存在")
		}
		return err
	}

	// 验证密码是否正确
	if user.PassWord != loginData.PassWord {
		return errors.New("用户密码错误")
	}

	// 验证通过，返回 nil
	return nil
}

// GetUserByID 根据用户ID获取用户信息
func GetUserByID(userID int) (models.User, error) {
	return models.GetUserByID(userID)
}
