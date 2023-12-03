package service

import (
	"onlineshop/models"
)

/**
 * @File : user.go
 * @Description : 用户相关的服务
 * @Author : chen
 * @Date : 2023/12/03
 */

func UsersRegister(theUser *models.User) (err error) {
	err = models.CreateAUser(theUser)

	return err
}