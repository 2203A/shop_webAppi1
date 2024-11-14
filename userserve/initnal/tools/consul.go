package tools

import (
	"SX1/shop_webAppi/userserve/cmd/global"
	"fmt"
	"github.com/hashicorp/consul/api"
	"strconv"
)

// consul结构体
type Consul struct {
	Host string
	Port int
}

// 实例化cunsul客户端
func NewConsulClient(c Consul) *Consul {
	return &Consul{
		Host: c.Host,
		Port: c.Port,
	}
}

// 类方法接口
type ConsulInterface interface {
	RegistrationServe(id string, name string, tags []string, port int, address string) error //注册服务
}

// 注册服务
func (c *Consul) RegistrationServe(id string, name string, tags []string, port int, address string) error {
	//初始化
	config := api.DefaultConfig()
	//设置consul访问地址
	config.Address = global.AppConfig.Consul.Host + ":" + strconv.Itoa(global.AppConfig.Consul.Port)
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("初始化客户端失败", err)
		return nil
	}
	//实例化注册服务结构体
	registration := api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Tags:    tags,
		Port:    port,
		Address: address,
	}
	//调用方法进行注册
	err = client.Agent().ServiceRegister(&registration)
	if err != nil {
		return nil
	}
	return err
}
