package entidy

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Mobile        string `gorm:"type:varchar(11);unique;not null;comment:手机号"`
	Password      string `gorm:"type:varchar(30);unique;not null;comment:密码"`
	Email         string `gorm:"type:varchar(30);unique;not null;comment:邮箱"`
	NickName      string `gorm:"type:varchar(30);unique;not null;comment:昵称"`
	Gender        int    `gorm:"type:int(2);unique;not null;comment:性别"`
	AvatarUrl     string `gorm:"type:varchar(30);unique;not null;comment:头像地址"`
	AccountStatus int    `gorm:"type:int(2);unique;not null;comment:账号状态"`
}
type Address struct {
	gorm.Model
	UserId         int    `gorm:"type:int(3);unique;not null;comment:用户id"`
	ConsigneeName  string `gorm:"type:varchar(30);not null;comment:地址用户名称"`
	ConsigneePhone string `gorm:"type:varchar(11);not null;comment:手机号"`
	Province       string `gorm:"type:varchar(11);not null;comment:省份"`
	City           string `gorm:"type:varchar(20);not null;comment:城市"`
	County         string `gorm:"type:varchar(20);not null;comment:县/区"`
	Detailed       string `gorm:"type:varchar(50);not null;comment:详细地址"`
	IsDefault      int    `gorm:"type:int(3);comment:是否是默认地址1是默认地址2不是"`
}
