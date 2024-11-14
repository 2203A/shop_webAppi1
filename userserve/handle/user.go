package handle

import (
	"SX1/shop_webAppi/userserve/initnal/entity"
	UserPDS "SX1/shop_webAppi/userserve/initnal/proto"
	"SX1/shop_webAppi/userserve/initnal/service"
	"context"
)

type User struct {
}

// 注册方法
func (u User) UserRegistrations(ctx context.Context, req *UserPDS.UserRegistrationsReq) (*UserPDS.UserRegistrationsResp, error) {
	user := entity.User{
		Phone:    req.Phone,
		Password: req.Password,
	}
	//调用注册方法
	registration, err := service.UserRegistration(user)
	if err != nil {
		return nil, err
	}
	res := UserPDS.UserRegistrationsResp{Id: int64(registration.ID)}
	return &res, nil
}
