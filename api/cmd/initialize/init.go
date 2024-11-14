package initialize

import (
	"SX1/shop_api/api/cmd/global"
	UserPDA "SX1/shop_api/api/initnal/proto/user"
	"SX1/shop_api/api/initnal/router"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/olivere/elastic/v7"
	"github.com/smartwalle/alipay/v3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	//初始化appconfig配置文件
	InitAppConfig()
	//初始化zap日志
	InitZap()
	//初始化路由组
	InitRouter()
	//初始化连接mysql
	InitMysql()
	//初始化支付宝客户端
	InitAlipay()
	//初始化连接用户服务
	InitUserServe()
	//初始化Redis
	InitRedis()
	//初始化Elasticsearch
	InitElasticsearch()
	//初始化连接rocketmq
	InitRocketmq()
}

//初始化连接用户服务

func InitUserServe() {
	// 1.连接
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常： %s\n", err)
	}
	// 2. 实例化gRPC客户端
	global.UserClient = UserPDA.NewUserClient(conn)
	fmt.Println("init user success")
}

//初始化路由组

func InitRouter() {
	g := gin.Default()
	router.Router(g)
	global.Router = g
}

//初始化zap日志

func InitZap() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{
		global.AppConfig.ZapConf.Path,
	}
	build, _ := config.Build()
	zap.ReplaceGlobals(build)
}

//初始化appconfig配置文件

func InitAppConfig() {
	viper.SetConfigFile("D:\\MyGo\\src\\SX1\\shop_webAppi\\api\\cmd\\appconfig.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("read viper err", err)
		return
	}
	err = viper.Unmarshal(&global.Nacos)
	if err != nil {
		fmt.Println("nacos err", err)
		return
	}
	fmt.Println("nacos success", global.Nacos.Address)
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         global.Nacos.Address, // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
		Username:            global.Nacos.User,
		Password:            global.Nacos.Pass,
	}
	// 至少一个ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: global.Nacos.Host,
			Port:   uint64(global.Nacos.Port),
		},
	}
	// 创建动态配置客户端
	client, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	config, err := client.GetConfig(vo.ConfigParam{
		DataId: global.Nacos.DataId,
		Group:  global.Nacos.Group,
	})
	if err != nil {
		fmt.Println("read nacos err", err)
		return
	}
	err = json.Unmarshal([]byte(config), &global.AppConfig)
	if err != nil {
		fmt.Println("appconfig赋值失败", err)
		return
	}
	fmt.Println("appconfig success", global.AppConfig.ZapConf.Path)
}

func InitMysql() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", global.AppConfig.MysqlConf.User, global.AppConfig.MysqlConf.Pass, global.AppConfig.MysqlConf.Host, global.AppConfig.MysqlConf.Port, "test")
	fmt.Println(dsn)
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("init mysql err", err)
		zap.S().Info("init mysql err", err)
		panic(err)
	}
	fmt.Println("init mysql success")
	zap.S().Info("init mysql success")
}

// 初始化Redis
func InitRedis() {
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     "118.89.80.34:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := global.RedisClient.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
	fmt.Println("init redis success")
}

//初始化支付宝客户端

func InitAlipay() {
	var err error
	global.AlipayClient, err = alipay.New(global.AppConfig.Alipay.AppID, global.AppConfig.Alipay.PrickKey, false)
	if err != nil {
		fmt.Println("init alipayClient err", err)
		zap.S().Info("init alipayClient err", err)
		return
	}
	err = global.AlipayClient.LoadAliPayPublicKey(global.AppConfig.Alipay.PublicKey)
	if err != nil {
		fmt.Println("load alipay publicKey err", err)
		zap.S().Info("load alipay publicKey err", err)
		return
	}
	fmt.Println("init alipay success")
}

//初始化Elasticsearch

func InitElasticsearch() {
	var err error
	global.ElasticClient, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://118.89.80.34:9200/"))
	if err != nil {
		fmt.Println("init elasticsearchClient err", err)
		zap.L().Info("init elasticsearchClient err" + err.Error())
		return
	}
	fmt.Println("init elasticsearchClient success")
	zap.S().Info("init elasticsearchClient success")
}

//初始化连接rocketmq

func InitRocketmq() {

}
