package models

import (
	"github.com/jinzhu/gorm"
	mysql "onlineshop/mysql"
)

/**
 * @File : comment.go
 * @Description : 评论模型 与其增删查改
 * @Author : chen
 * @Date : 2023/11/16
 */

type Comment struct {
	gorm.Model
	UserID    int `json:"userId"`
	ProductID int `json:"productId"`
	Likes     int `json:"likes"`
}

// 创建 Comment
func CreateAComment(theComment *Comment) (err error) {
	err = mysql.DB.Create(&theComment).Error

	return err
}

// 获取所有 Comment
func GetAllComment() (theCommentList []*Comment, err error) {
	err = mysql.DB.Find(&theCommentList).Error

	if err != nil {
		return nil, err
	}

	return theCommentList, nil
}

// 通过 id 获取 Comment
func GetCommentByID(id int) (theComment *Comment, err error) {
	err = mysql.DB.Where("id = ?", id).First(&theComment).Error

	if err != nil {
		return nil, err
	}

	return theComment, nil
}

// 更新一个存在于数据库的 Comment
func UpdateAComment(theComment *Comment) (err error) {
	err = mysql.DB.Save(&theComment).Error

	return err
}

// 通过 id 删除 Comment
func DeleteCommentByID(id int) (err error) {
	err = mysql.DB.Where("id = ?", id).Delete(Comment{}).Error

	return err
}
