package global

import (
	"github.com/gin-gonic/gin"
	"shop-APi/intelnal/proto/UserSrv"
	"shop-APi/pkg/appconfig"
)

var (
	Router        *gin.Engine
	UserSrvClient UserSrv.UserClient
	InitConfig    appconfig.AppConfig
)
