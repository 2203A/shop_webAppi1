package main

import (
	"context"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shop-APi/cmd/global"
	_ "shop-APi/pkg/initialization"
	"time"
)

func main() {
	zap.S().Infow(
		"Started ApiServer：：：",
		zap.String("Host", "127.0.0.1"),
		zap.Int("Port", 8080),
	)

	zap.S().Info("Api is running!!!!")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: global.Router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
