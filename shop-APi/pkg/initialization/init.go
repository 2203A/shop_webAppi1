package initialization

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"shop-APi/cmd/global"
	"shop-APi/intelnal/proto/UserSrv"
	routerGroup "shop-APi/intelnal/router"
	"shop-APi/pkg/appconfig"
)

func init() {
	//初始化nacos
	InitNacos()
	//初始化用户服务端的链接
	InitUserSrv()
	//初始化gin框架
	InitRouter()
}

func InitUserSrv() {
	// 1.连接
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常： %s\n", err)
	}
	global.UserSrvClient = UserSrv.NewUserClient(conn)
}

func InitNacos() {

	//create clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "5ae01a4a-b833-43e4-a0fd-57c372442648", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
		Username:            "nacos",
		Password:            "admin",
	}
	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "121.36.211.57",
			Port:   8848,
		},
	}
	// Create config client for dynamic configuration
	configClient, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "xzcUserApi",
		Group:  "dev"})
	if err != nil {
		zap.S().Errorw("get config failed", zap.Error(err))
	}

	global.InitConfig = appconfig.AppConfig{}
	json.Unmarshal([]byte(content), &global.InitConfig)
	fmt.Println(global.InitConfig)

}

func InitRouter() {
	r := gin.Default()
	routerGroup.RouterGroup(r)
	global.Router = r

}
