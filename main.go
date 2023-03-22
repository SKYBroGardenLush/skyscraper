package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"skyscraper/framework"
	"skyscraper/framework/middleware"
	"skyscraper/provider/demo"
	"skyscraper/test"
	"syscall"
	"time"
)

func main() {
	// 创建core结构
	core := framework.NewCore()
	// 绑定具体的服务
	core.Bind(&demo.DemoServiceProvider{})

	core.Use(middleware.Cost())

	test.RegisterRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}

	// 这个goroutine是启动服务的goroutine
	go func() {
		server.ListenAndServe()
	}()

	// 当前的goroutine等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIcoreT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前goroutine等待信号
	<-quit

	// 调用Server.Shutdown graceful结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
