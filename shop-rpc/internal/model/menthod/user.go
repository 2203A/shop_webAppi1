package menthod

import (
	"examProject/cmd/userSrv/proto/UserSrv"
	"examProject/internal/model/entidy"
	"examProject/pkg/global"
)

func UserLogin(user *UserSrv.UserRegisterReq) (entidy.User, error) {
	var stu entidy.User
	err := global.Db.Where("mobile=?", user.Mobile).First(&stu).Error
	return stu, err
}
func UserRegister(user entidy.User) (entidy.User, error) {
	err := global.Db.Create(&user).Error
	return user, err
}
