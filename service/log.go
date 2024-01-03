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

// GetLogList 获取日志列表服务函数
func GetLogList() ([]map[string]interface{}, error) {
	// 调用数据库模型的方法获取日志列表
	logs, err := models.GetAllLog()
	if err != nil {
		return nil, err
	}

	// 构造返回数据
	var logList []map[string]interface{}
	for _, log := range logs {
		logData := map[string]interface{}{
			"Id":         log.ID,
			"submitName": log.UserID, // 替换为实际的用户字段
			"time_data":  log.UpdatedAt.Format("2006/01/02 15:04:05"),
			"title":      log.Title,
			"text":       log.Content,
			// 添加其他需要返回的日志信息字段
		}
		logList = append(logList, logData)
	}

	return logList, nil
}
