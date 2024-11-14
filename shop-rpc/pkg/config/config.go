package config

type ApiConfig struct {
	Server
	Log
	MysqlMaster
	Consul
}

type Server struct {
	Host string   `mapstructure:"host"`
	Port int      `mapstructure:"port"`
	Name string   `mapstructure:"name"`
	Tags []string `mapstructure:"tags"`
}

type Log struct {
	LogPath string `mapstructure:"logPath"`
}

type MysqlMaster struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"dbName"`
}

type Consul struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
