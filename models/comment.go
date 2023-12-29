package models

import (
	"github.com/jinzhu/gorm"
	mysql "onlineshop/mysql"
	"time"
)

/**
 * @File : comment.go
 * @Description : 评论模型 与其增删查改
 * @Author : chen
 * @Date : 2023/11/16
 */

type Comment struct {
	gorm.Model
	UserID    int       `json:"userId" gorm:"foreignKey:UserID"` // 保留作为外键
	ProductID int       `json:"productId" gorm:"foreignKey:ProductID"`
	User      User      `gorm:"foreignKey:UserID"` // 新增，定义关联的用户
	Likes     int       `json:"likes"`
	Content   string    `json:"content"`
	Star      int       `json:"star"`
	ParentID  int       `json:"parentId"`          // 指向父评论的ID
	Replies   []Comment `json:"replies,omitempty"` // 用来存储回复
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
func GetCommentByID(id int) (theComment Comment, err error) {
	err = mysql.DB.Where("id = ?", id).First(&theComment).Error
	if err != nil {
		return theComment, err
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

// 通过产品ID获取评论及其回复
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04") // Go 使用这种看起来奇怪的日期作为格式化模板
}

type CommentResponse struct {
	CommentID uint              `json:"commentId"`
	UserID    uint              `json:"userId"`
	Username  string            `json:"username"`
	Content   string            `json:"content"`
	Time      string            `json:"time"`
	GoodNum   int               `json:"goodNum"`
	Star      int               `json:"star,omitempty"` // 使用omitempty，如果Star为0则不显示
	Replies   []CommentResponse `json:"reply,omitempty"`
}

// 构造单个评论响应
func NewCommentResponse(c Comment, isReply bool) CommentResponse {
	resp := CommentResponse{
		CommentID: c.ID,
		UserID:    c.User.ID,
		Username:  c.User.UserName,
		Content:   c.Content,
		Time:      FormatTime(c.CreatedAt),
		GoodNum:   c.Likes,
	}

	if !isReply {
		resp.Star = c.Star // 只有非回复（即主评论）才设置Star值
	}

	return resp
}

func GetCommentsAndRepliesByProductID(productId int) ([]CommentResponse, error) {
	var comments []Comment
	var responses []CommentResponse

	err := mysql.DB.Where("product_id = ? AND parent_id = 0", productId).Preload("User").Find(&comments).Error
	if err != nil {
		return nil, err
	}

	for _, comment := range comments {
		commentResponse := NewCommentResponse(comment, false) // 这是主评论

		var replies []Comment
		err := mysql.DB.Where("parent_id = ?", comment.ID).Preload("User").Find(&replies).Error
		if err != nil {
			continue
		}

		for _, reply := range replies {
			replyResponse := NewCommentResponse(reply, true) // 这是回复
			commentResponse.Replies = append(commentResponse.Replies, replyResponse)
		}

		responses = append(responses, commentResponse)
	}

	return responses, nil
}

// 用户发表评论
type AddCommentRequest struct {
	UserID  int    `json:"userId"`
	Content string `json:"content"`
}
