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
