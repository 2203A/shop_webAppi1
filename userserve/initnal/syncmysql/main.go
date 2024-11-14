package main

import (
	"SX1/shop_webAppi/userserve/cmd/global"
	"SX1/shop_webAppi/userserve/initnal/entity"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var err error
	dsn := "root:D70E2C5F2FB08C03@tcp(118.89.80.34:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("init mysql err", err)
		zap.S().Info("init mysql err", err)
		panic(err)
	}
	fmt.Println("init mysql success")
	global.DB.AutoMigrate(&entity.User{})
	zap.S().Info("init mysql success")
}
