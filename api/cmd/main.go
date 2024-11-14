package main

import (
	"SX1/shop_api/api/cmd/global"
	_ "SX1/shop_api/api/cmd/initialize"
	"go.uber.org/zap"
)

func main() {
	zap.L().Info("Init ApiServe Success")
	global.Router.Run()
}
