package global

import (
	"examProject/pkg/config"
	"gorm.io/gorm"
)

var (
	ApiConfig *config.ApiConfig //Api端的配置项
	Db        *gorm.DB          //全局的数据库操作
)
