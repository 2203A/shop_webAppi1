package initialization

import (
	"examProject/internal/model/entidy"
	"examProject/pkg/global"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {

	// 数据库配置
	InitMysql()
}

func InitMysql() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "mysql_DAJCYW", "175.24.164.107", 3306, "shopProject")
	global.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("mysql数据库初始化成功")
	global.Db.AutoMigrate(new(entidy.User), entidy.Address{})
}
