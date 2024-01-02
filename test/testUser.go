package test

import (
	"onlineshop/models"
)

var testUser = []models.User{
	{
		UserName: "黑",
		PassWord: "123",
		PhoneNum: "133123",
		UserType: 1,
		Email:    "123@123.com",
	},
	{
		UserName: "红",
		PassWord: "123456",
		PhoneNum: "133789",
		UserType: 1,
		Email:    "456@123.com",
	},
	{
		UserName: "白",
		PassWord: "123789",
		PhoneNum: "133456",
		UserType: 1,
		Email:    "789@123.com",
	},
}

func addUser() {
	for _, i := range testUser {
		models.CreateAUser(&i)
	}
}