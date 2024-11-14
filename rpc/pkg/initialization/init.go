package initialization

import (
	"examProject/pkg/golbal"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	// 读取配置
	InitConfig()
	//配置日志
	InitZap()
	// 数据库配置
	InitMysql()
}

func InitMysql() {
	mysqlConfig := golbal.ApiConfig.MysqlMaster
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DbName)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)
	// 全局模式
	var err error
	golbal.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		zap.S().Infow("mysql connetion failed !", err)
	}
	zap.S().Infow("数据库初始化成功", zap.Int("code", http.StatusOK))
}

func InitConfig() {
	// 读取配置文件
	v := viper.New()
	//设置配置文件路径
	v.SetConfigFile("/Users/fuyx/GolandProjects/zg4ProjectServer/application-dev.yaml")
	//读取配置文件
	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	// 进行结构体映射并设置全局变量 -- 序列化
	if err := v.Unmarshal(&golbal.ApiConfig); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}
}

func InitZap() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{
		golbal.ApiConfig.Log.LogPath,
		"stdout",
	}
	logger, _ := config.Build()
	//全局变量配置
	zap.ReplaceGlobals(logger)
}
