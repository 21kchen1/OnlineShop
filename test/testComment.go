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

var testCommentList = []models.Comment{
	{
		UserID:    1,
		ProductID: 1,
		Content:   "好好好",
	},
	{
		UserID:    1,
		ProductID: 1,
		Content:   "真不错",
	},
	{
		UserID:    1,
		ProductID: 1,
		Content:   "挺合适",
	},
}

func addComment() {
	for _, i := range testCommentList {
		models.CreateAComment(&i)
	}
}
