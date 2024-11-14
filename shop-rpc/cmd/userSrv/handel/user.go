package handel

import (
	"context"
	"examProject/cmd/userSrv/proto/UserSrv"
	"examProject/internal/model/entidy"
	"examProject/internal/model/menthod"
	"fmt"
)

type User struct {
}

func (u *User) UserAddressAdd(ctx context.Context, req *UserSrv.UserAddressAddReq) (*UserSrv.UserAddressAddRes, error) {
	//TODO implement me
	panic("implement me")
}

func (u *User) UserRegister(ctx context.Context, req *UserSrv.UserRegisterReq) (*UserSrv.UserRegisterRes, error) {
	fmt.Println(req)
	userlogin, _ := menthod.UserLogin(req)
	var Id = 0
	if userlogin.ID == 0 {
		user := entidy.User{
			Mobile:   req.Mobile,
			Password: req.Password,
		}
		userRegister, err := menthod.UserRegister(user)
		if err != nil {
			return nil, fmt.Errorf("注册失败")
		}
		Id = int(userRegister.ID)
	} else {
		if userlogin.Password != req.Password {
			return nil, fmt.Errorf("密码输入错误")
		} else {
			Id = int(userlogin.ID)
		}
	}
	return &UserSrv.UserRegisterRes{Id: int64(Id)}, nil

}
