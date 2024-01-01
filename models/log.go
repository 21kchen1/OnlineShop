package models

import (
	"github.com/jinzhu/gorm"
	mysql "onlineshop/mysql"
)

/**
 * @File : log.go
 * @Description : 日志模型 与其增删查改
 * @Author : chen
 * @Date : 2023/11/17
 */

type Log struct {
	gorm.Model
	UserID  int    `json:"userId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// 创建 Log
func CreateALog(theLog *Log) (err error) {
	err = mysql.DB.Create(&theLog).Error

	return err
}

// 获取所有 Log
func GetAllLog() (theLogList []*Log, err error) {
	err = mysql.DB.Find(&theLogList).Error

	if err != nil {
		return nil, err
	}

	return theLogList, nil
}

// 通过 id 获取 Log
func GetLogByID(id int) (theLog Log, err error) {
	err = mysql.DB.Where("id = ?", id).First(&theLog).Error

	if err != nil {
		return theLog, err
	}

	return theLog, nil
}

// 更新一个存在于数据库的 Log
func UpdateALog(theLog *Log) (err error) {
	err = mysql.DB.Save(&theLog).Error

	return err
}

// 通过 id 删除 Log
func DeleteLogByID(id int) (err error) {
	err = mysql.DB.Where("id = ?", id).Delete(Log{}).Error

	return err
}
