package client

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbClient *gorm.DB
var errClient error

func init() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:dmy123123@tcp(127.0.0.1:3306)/mydb?charset=utf8&parseTime=True&loc=Local", // username:password@(ip:port)+db_name+ DSN data source name
		DontSupportRenameColumn:   true,                                                                            // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                           // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	DbClient = db
	errClient = err
	fmt.Println(err)
}
