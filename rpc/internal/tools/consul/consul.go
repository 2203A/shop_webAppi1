package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

var (
	consulConfig *api.Config
)

/**
@desc consul配置项
@host 指consul服务的host
@port 指consul服务的port
*/

type Consul struct {
	Host string
	Port int
}

type ConsulClient interface {
	Register(host string, port int, name string, tag []string, id string) // host-> 要被注册的服务的host port->要被注册的服务的port
}

func NewConsulClient(host string, port int) ConsulClient {
	consulConfig = api.DefaultConfig()
	return &Consul{
		Host: host,
		Port: port,
	}
}

/**
@author fuyx
@date 2041/08/28
@desc  注册到 consul
*/

func (consul *Consul) Register(host string, port int, name string, tag []string, id string) {
	consulConfig.Address = fmt.Sprintf("%s:%d", consul.Host, consul.Port) //consul 地址  127.0.0.1:8500
	consuleClient, err := api.NewClient(consulConfig)
	if err != nil {
		zap.S().Error(err)
	}
	registeration := new(api.AgentServiceRegistration)
	registeration.ID = id
	registeration.Name = name
	registeration.Tags = tag
	registeration.Address = host
	registeration.Port = port
	registerError := consuleClient.Agent().ServiceRegister(registeration)
	if registerError != nil {
		zap.S().Error(registerError)
	}
}

func GetServiceList() map[string]*api.AgentService {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	consuleClient, err := api.NewClient(config)
	if err != nil {
		zap.S().Error(err)
	}
	servicesList, err := consuleClient.Agent().Services()
	if err != nil {
		zap.S().Error(err)
	}
	return servicesList
}
