package test

import (
	"onlineshop/models"
)

/**
 * @File : testFavorites.go
 * @Description : 评论相关功能测试
 * @Author : chen
 * @Date : 2023/12/27
 */

var testCommentList = []models.Comment {
	{
		UserID: 123,
	},
	{
		UserID: 1213,
	},
	{
		UserID: 1243,
	},
}

func AddComment() {
	for _, i := range testCommentList {
		models.CreateAComment(&i)
	}
}