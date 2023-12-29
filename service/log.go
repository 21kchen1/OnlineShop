package service

import (
	"fmt"
	"onlineshop/models"
)

// AddLog 添加日志服务函数
func AddLog(userID int, title, content string) error {
	// 创建日志对象
	log := models.Log{
		UserID:  userID,
		Title:   title,
		Content: content,
	}

	// 调用数据库模型的方法添加日志
	err := models.CreateALog(&log)
	if err != nil {
		return fmt.Errorf("添加日志失败：%w", err)
	}

	return nil
}

// DeleteLog 删除日志服务函数
func DeleteLog(logID int) error {
	// 调用数据库模型的方法删除日志
	err := models.DeleteLogByID(logID)
	if err != nil {
		return fmt.Errorf("删除日志失败：%w", err)
	}

	return nil
}

// EditLog 修改日志服务函数
func EditLog(logID int, title, content string) error {
	// 获取日志对象
	log, err := models.GetLogByID(logID)
	if err != nil {
		return fmt.Errorf("获取日志失败：%w", err)
	}

	// 更新日志信息
	log.Title = title
	log.Content = content

	// 调用数据库模型的方法保存更新后的日志
	err = models.UpdateALog(&log)
	if err != nil {
		return fmt.Errorf("修改日志失败：%w", err)
	}

	return nil
}
