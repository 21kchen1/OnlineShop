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