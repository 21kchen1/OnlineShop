
package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"onlineshop/models"
	"onlineshop/service"
)

/**
 * @File : user.go
 * @Description : 用户相关的控制器
 * @Author : chen
 * @Date : 2023/12/03
 */

// 注册模块
func UsersRegister(c *gin.Context) {
	var theUser models.User

	// 自动获取数据
	if err := c.ShouldBind(&theUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "c.ShouldBind 获取参数失败",
		})
		return
	}

	if err := service.CheckUserExistByName(&theUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "CheckUserExistByName 用户名重复",
		})
		return
	}

	if err := service.CreateUsers(&theUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "service.UsersRegister 创建失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "UsersRegister 用户创建成功",
	})
}

// 用户登录模块
func UsersLogin(c *gin.Context) {
	fmt.Println("Entering UsersLogin")

	var loginData models.User

	// 手动获取参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	loginData.UserName = username
	loginData.PassWord = password
	fmt.Println("UsersLogin - Username:", username, "Password:", password)
	// 调用验证函数
	if err := service.CheckUserLogin(&loginData); err != nil {
		fmt.Println("Error in CheckUserLogin:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       err.Error(),
		})
		return
	}

	// 获取用户信息
	user, err := models.GetUserByName(loginData.UserName)
	if err != nil {
		fmt.Println("Error in GetUserByName:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "获取用户信息失败",
		})
		return
	}

	// 验证成功，返回相应信息
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "UsersLogin 用户登录成功",
		"userId":    user.ID,
	})
}

/**
 * @File : user.go
 * @Description : 用户管理
 * @Author : you
 * @Date : 2023/12/03
 */

// 更新用户信息

// 根据 role 映射到 UserType 的辅助函数
func getUserTypeFromRole(role string) int {
	switch role {
	case "普通用户":
		return 0
	case "商家":
		return 1
	case "管理员":
		return 2
	default:
		// 默认返回普通用户类型
		return 0
	}
}

// 在 getRoleFromUserType 函数中添加对 0、1 和 2 的映射
func getRoleFromUserType(userType int) string {
	switch userType {
	case 0:
		return "普通用户"
	case 1:
		return "商家"
	case 2:
		return "管理员"
	default:
		// 默认返回普通用户角色
		return "普通用户"
	}
}

// 更新用户信息
func UpdateUserInfo(c *gin.Context) {
	var requestData struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Role     string `json:"role"`
		Password string `json:"password"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "获取参数失败",
		})
		return
	}

	// 创建一个 User 对象并设置需要更新的字段，将 Role 映射到 UserType
	updatedUser := &models.User{
		Model:    gorm.Model{ID: requestData.ID},
		UserName: requestData.Username,
		Email:    requestData.Email,
		PhoneNum: requestData.Phone,
		UserType: getUserTypeFromRole(requestData.Role),
		PassWord: requestData.Password, // 添加密码字段处理
	}

	// 调用 models 中的更新用户信息函数
	if err := models.UpdateAUser(updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "更新用户信息失败",
		})
		return
	}

	// 返回更新后的用户信息
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "更新用户信息成功",
	})
}

// 添加用户
func AddUser(c *gin.Context) {
	var requestData struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Phone    string `json:"phone"`
		Role     string `json:"role"`
		Password string `json:"password" binding:"required"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "获取参数失败",
		})
		return
	}

	// 创建一个 User 对象并设置字段
	newUser := &models.User{
		UserName: requestData.Username,
		Email:    requestData.Email,
		PhoneNum: requestData.Phone,
		UserType: getUserTypeFromRole(requestData.Role),
		PassWord: requestData.Password,
	}

	// 调用 models 中的添加用户函数
	if err := models.CreateAUser(newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "添加用户失败",
		})
		return
	}

	// 返回添加后的用户信息
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "添加用户成功",
	})
}

// 获取用户列表
func GetUserList(c *gin.Context) {
	// 调用 models 中的获取所有用户函数
	userList, err := models.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "获取用户列表失败",
		})
		return
	}

	// 构建只包含指定信息的用户列表
	simplifiedUserList := make([]gin.H, 0)
	for _, user := range userList {
		simplifiedUserList = append(simplifiedUserList, gin.H{
			"id":       user.ID,
			"username": user.UserName,
			"email":    user.Email,
			"phone":    user.PhoneNum,
			"role":     getRoleFromUserType(user.UserType),
		})
	}

	// 返回用户列表
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "获取用户列表成功",
		"data":      simplifiedUserList,
	})
}

// 删除用户

func DeleteUser(c *gin.Context) {
	var requestData struct {
		UserID int `json:"userID" binding:"required"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "获取参数失败",
		})
		return
	}

	// 调用 models 中的通过 ID 删除用户函数
	err := models.DeleteUserByID(requestData.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "删除用户失败",
		})
		return
	}

	// 获取删除后的用户列表
	userList, err := models.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "获取用户列表失败",
		})
		return
	}

	// 构建只包含指定信息的用户列表
	simplifiedUserList := make([]gin.H, 0)
	for _, user := range userList {
		simplifiedUserList = append(simplifiedUserList, gin.H{
			"id":       user.ID,
			"username": user.UserName,
			"email":    user.Email,
			"phone":    user.PhoneNum,
			"role":     getRoleFromUserType(user.UserType),
		})
	}

	// 返回删除后的用户列表
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "删除用户成功",
	})
}

/**
 * @File : user.go
 * @Description : 商家管理
 * @Author : you
 * @Date : 2023/12/29
 */

// 获取商家列表

func GetSellerList(c *gin.Context) {
	// 调用 models 中的获取商家列表函数
	sellerList, err := models.GetSellerList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "获取商家列表失败",
		})
		return
	}

	simplifiedSellerList := make([]gin.H, 0)
	for _, seller := range sellerList {
		simplifiedSellerList = append(simplifiedSellerList, gin.H{
			"id":         seller.ID,
			"sellername": seller.UserName,
			"mail":       seller.Email,
			"phone":      seller.PhoneNum,
			"address":    seller.Seller.Address,
		})
	}

	// 返回商家列表
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "获取商家列表成功",
		"data":      simplifiedSellerList,
	})
}

func AddSeller(c *gin.Context) {
	var requestData struct {
		SellerName string `json:"sellername" binding:"required"`
		Email      string `json:"mail" binding:"required"`
		Phone      string `json:"phone"`
		Password   string `json:"password" binding:"required"`
		Address    string `json:"address"` // 确保包含地址
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "获取参数失败",
		})
		return
	}

	// 创建一个 User 对象并设置字段
	newSeller := &models.User{
		UserName: requestData.SellerName,
		Email:    requestData.Email,
		PhoneNum: requestData.Phone,
		UserType: 1, // 商家用户类型
		PassWord: requestData.Password,
		Seller: models.Seller{ // 创建关联的Seller信息
			Address: requestData.Address,
		},
	}

	// 调用 models 中的添加用户函数
	if err := models.CreateAUser(newSeller); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "添加商家失败",
		})
		return
	}

	// 返回添加后的商家信息
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "添加商家成功",
	})
}

// 删除商家

func DeleteSeller(c *gin.Context) {
	var requestData struct {
		SellerID int `json:"sellerID" binding:"required"`
	}

	// 获取请求参数
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "获取参数失败",
		})
		return
	}

	// 调用 models 中的通过 ID 删除商家函数
	err := models.DeleteSellerByID(requestData.SellerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "删除商家失败",
		})
		return
	}

	// 获取删除后的商家列表
	sellerList, err := models.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "获取商家列表失败",
		})
		return
	}

	// 构建只包含指定信息的商家列表
	simplifiedSellerList := make([]gin.H, 0)
	for _, seller := range sellerList {
		if seller.UserType == 1 { // 商家用户类型
			simplifiedSellerList = append(simplifiedSellerList, gin.H{
				"id":         seller.ID,
				"sellername": seller.UserName,
				"mail":       seller.Email,
				"phone":      seller.PhoneNum,
			})
		}
	}

	// 返回删除后的商家列表
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "删除商家成功",
	})
}

//更新商家信息

type UpdateSellerRequest struct {
	ID         uint   `json:"id"`
	SellerName string `json:"sellername"` // 对应前端的sellername字段
	Email      string `json:"mail"`       // 对应前端的mail字段
	Phone      string `json:"phone"`      // 对应前端的phone字段
	Password   string `json:"password"`   // 对应前端的password字段
	Address    string `json:"address"`    // 对应前端的address字段
}

func UpdateSellerInfo(c *gin.Context) {
	var requestData UpdateSellerRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"isSuccess": false,
			"msg":       "获取参数失败",
		})
		return
	}

	// 获取并更新商家用户信息
	user, err := models.GetUserByID(int(requestData.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "找不到指定的商家",
		})
		return
	}

	// 更新用户基本信息
	user.UserName = requestData.SellerName
	user.Email = requestData.Email
	user.PhoneNum = requestData.Phone
	user.PassWord = requestData.Password
	user.Seller.Address = requestData.Address

	// 调用 models 中的更新用户信息函数
	if err := models.UpdateAUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"isSuccess": false,
			"msg":       "更新商家信息失败",
		})
		return
	}

	// 返回更新后的商家信息
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
		"msg":       "更新商家信息成功",
	})
}
