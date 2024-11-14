package userSrv

import (
	"context"
	"shop-APi/cmd/global"
	"shop-APi/intelnal/formValidate/userReq"
	"shop-APi/intelnal/proto/UserSrv"
)

func UserRegister(user userReq.UserRegister) (*UserSrv.UserRegisterRes, error) {
	register, err := global.UserSrvClient.UserRegister(context.Background(), &UserSrv.UserRegisterReq{
		Mobile:   user.Mobile,
		Password: user.Password,
	})
	return register, err
}
