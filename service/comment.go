package service

import (
	"onlineshop/models"
)

// DeleteComment 删除评论服务函数
func DeleteComment(commentID int) error {
	// 调用数据库模型的方法删除评论
	err := models.DeleteCommentByID(commentID)
	if err != nil {
		return err
	}

	return nil
}

/**
 * @File : comment.go
 * @Description : 获取所有评论
 * @Author : chen
 * @Date : 2024-1-1
 */
func GetAllComment() (commentList []*models.Comment, err error) {
	commentList, err = models.GetAllComment()
	if err != nil {
		return
	}
	return
}