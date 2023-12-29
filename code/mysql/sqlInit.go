package mysql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/**
 * @File : sqlInit.go
 * @Description : 数据库连接，创建表，提供 DB 用于其他类调用
 * @Author : chen
 * @Date : 2023/11/10
 */

var (
	DB *gorm.DB
)

func SqlInit() (err error) {
	userName := "ecommerce"        // 账户
	passWord := "hbfejftWLDs8s3pM" // 密码
	host := "45.153.131.189"       // 数据库地址
	port := 3306                   // 数据库端口
	DBName := "ecommerce"          // 数据库名
	timeOut := "2s"                // 超时限制

	// 拼接 dsm [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", userName, passWord, host, port, DBName, timeOut)

	DB, err = gorm.Open("mysql", dsn)

	return err
}

func Close() {
	DB.Close()
}
