package tools

// rocketmq结构体
type Rocketmq struct {
	Host string
	Port int
}

// 创建rocket客户端
func NewRocketmqClient(r Rocketmq) *Rocketmq {
	return &Rocketmq{
		Host: r.Host,
		Port: r.Port,
	}
}

// rocketmq类方法接口
type RocketmqInterface interface {
}
