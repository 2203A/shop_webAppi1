package userService

import (
	"SX1/shop_api/api/cmd/global"
	"SX1/shop_api/api/initnal/formvalidata"
	UserPDA "SX1/shop_api/api/initnal/proto/user"
	"context"
)

func UserRegistration(user formvalidata.User) (*UserPDA.UserRegistrationsResp, error) {
	registrations, err := global.UserClient.UserRegistrations(context.Background(), &UserPDA.UserRegistrationsReq{
		Phone:    user.Phone,
		Password: user.Password,
	})
	return registrations, err
}
