package appconfig

type AppConfig struct {
	Jwt       `json:"Jwt"`
	Consul    `json:"Consul"`
	MysqlConf `json:"MysqlConf"`
	ZapConf   `json:"ZapConf"`
	Alipay    `json:"Alipay"`
	RedisConf `json:"RedisConf"`
}

type Nacos struct {
	Address string
	Host    string
	Port    int
	User    string
	Pass    string
	DataId  string
	Group   string
}

type Jwt struct {
	SecretKey  string `json:"secretKey"`
	ExpireTime int    `json:"expireTime"`
}

type Consul struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type MysqlConf struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ZapConf struct {
	Path string `json:"path"`
}

type Alipay struct {
	AppID     string `json:"appID"`
	PublicKey string `json:"publicKey"`
	PrickKey  string `json:"prickKey"`
}
type RedisConf struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
