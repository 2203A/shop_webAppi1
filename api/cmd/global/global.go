package global

import (
	"SX1/shop_api/api/cmd/appconfig"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/olivere/elastic/v7"
	"github.com/smartwalle/alipay/v3"
	"gorm.io/gorm"
)

var (
	Nacos         *appconfig.Nacos
	AppConfig     *appconfig.AppConfig
	Router        *gin.Engine
	AlipayClient  *alipay.Client
	DB            *gorm.DB
	RedisClient   *redis.Client
	ElasticClient *elastic.Client
)
