package run

import (
	"fmt"
	"log"
	"myWebCli/settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/context"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ForeverElegant(e *gin.Engine) {
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", settings.Conf.App.Host, settings.Conf.App.Port),
		Handler: e,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅的关闭服务器，为关闭服务器操作设置一个5s的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的CTRL+C就是出发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify 把收到的 syscall.SIGINT或syscall.SIGTERM 信号发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会堵塞
	<-quit                                               // 堵塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Debug("ShutDown Server....")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server ShutDown: ", zap.Error(err))
	}
	zap.L().Debug("Server exiting")
}
