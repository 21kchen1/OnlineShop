/**
 * @File : comment.go
 * @Description : 评论管理的控制器
 * @Author : you
 * @Date : 2023-12-28
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlineshop/models"
	"onlineshop/service"
	"onlineshop/util"
)

/**
 * @File : comment.go
 * @Description : 处理获取商品评论的请求
 * @Author : you
 * @Date : 2023-12-28
 */

type ProductCommentRequest struct {
	ProductID int `json:"productId"`
}

func GetCommentsByProductID(c *gin.Context) {
	var req ProductCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comments, err := models.GetCommentsAndRepliesByProductID(req.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// 发表评论

type AddCommentRequest struct {
	UserID    int    `json:"userId"`
	ProductID int    `json:"productId"`
	Content   string `json:"content"`
	Star      int    `json:"star"`
}

func AddComment(c *gin.Context) {
	var req AddCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"isSuccess": false, "msg": err.Error()})
		return
	}

	// 验证星级评分是否在0到5之间
	if req.Star < 0 || req.Star > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"isSuccess": false, "msg": "Star rating must be between 0 and 5"})
		return
	}

	// 创建一个新的评论实例
	newComment := models.Comment{
		UserID:    req.UserID,
		ProductID: req.ProductID,
		Content:   req.Content,
		ParentID:  0,        // 0表示这是一个主评论
		Likes:     0,        // 新评论的点赞数初始为0
		Star:      req.Star, // 使用请求中提供的星级评分
	}

	// 调用模型中的函数来保存新评论
	err := models.CreateAComment(&newComment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"isSuccess": false, "msg": "Failed to add comment"})
	} else {
		c.JSON(http.StatusOK, gin.H{"isSuccess": true, "msg": "Comment added successfully"})
	}
}

/**
 * @File : comment.go
 * @Description : 用户回复评论
 * @Author : you
 * @Date : 2023-12-28
 */

type AddReplyRequest struct {
	UserID    int    `json:"userId"`
	CommentID int    `json:"commentId"`
	Content   string `json:"content"`
}

func AddReply(c *gin.Context) {
	var req AddReplyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"isSuccess": false, "msg": err.Error()})
		return
	}

	// 创建一个新的回复实例
	newReply := models.Comment{
		UserID:    req.UserID,
		ProductID: 1,             // TODO: 设置为被回复的评论所属的产品ID
		ParentID:  req.CommentID, // 设置为被回复的评论的ID
		Content:   req.Content,
		Likes:     0, // 新回复的点赞数初始为0
		Star:      0, // 回复通常不需要星级评分
	}

	// 调用模型中的函数来保存新回复
	err := models.CreateAComment(&newReply) // 可以复用创建评论的逻辑
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"isSuccess": false, "msg": "Failed to add reply"})
	} else {
		c.JSON(http.StatusOK, gin.H{"isSuccess": true, "msg": "Reply added successfully"})
	}
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	var requestData struct {
		CommentID int `json:"commentId"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		util.ErrRespon(c, err, "获取参数失败")
		return
	}

	// 调用 service 层删除评论的函数
	err := service.DeleteComment(requestData.CommentID)

	if err != nil {
		util.ErrRespon(c, err, "删除评论失败")
		return
	}

	// 返回成功的响应
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "删除评论成功",
	})
}

/**
 * @File : comment.go
 * @Description : 获取所有评论
 * @Author : chen
 * @Date : 2024-1-1
 */
func GetAllComment(c *gin.Context) {
	commentList, err := service.GetAllComment()

	if err != nil {
		util.ErrRespon(c, err, "获取所有评论失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "获取所有评论成功",
		"data":      commentList,
	})
}
