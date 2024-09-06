package models

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// https://gorm.io/zh_CN/docs/connecting_to_the_database.html

var DB *gorm.DB

func init() {
	// 加载 .ini 文件，连接数据库
	cfg, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	ip := cfg.Section("mysql").Key("ip").String()
	port := cfg.Section("mysql").Key("port").String()
	user := cfg.Section("mysql").Key("user").String()
	password := cfg.Section("mysql").Key("password").String()
	database := cfg.Section("mysql").Key("database").String()

	// 与数据库建立连接
	// dsn := "root:scx716@tcp(127.0.0.1:3306)/shop?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 打印 sql
		QueryFields: true,
		// 禁用事务
		// SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println(err)
	}
}
