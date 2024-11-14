package service

import (
	"SX1/shop_webAppi/userserve/cmd/global"
	"SX1/shop_webAppi/userserve/initnal/entity"
)

// 注册方法
func UserRegistration(user entity.User) (entity.User, error) {
	err := global.DB.Create(&user).Error
	return user, err
}
