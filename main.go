package main

import (
	"skyscraper/app/console"
	"skyscraper/app/http"
	"github.com/SKYBroGardenLush/skyscraper/framework"
	"github.com/SKYBroGardenLush/skyscraper/framework/containerService/service/app"
	"github.com/SKYBroGardenLush/skyscraper/framework/containerService/service/kernel"
)

func main() {
	// 初始化服务容器
	container := framework.NewSkyscraperContainer()
	// 绑定App服务提供者
	container.Bind(&app.AppProvider{})
	// 后续初始化需要绑定的服务提供者...

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpCore(); err == nil {
		container.Bind(&kernel.KernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)

	//// 这个goroutine是启动服务的goroutine
	//go func() {
	//	server.ListenAndServe()
	//}()
	//
	//// 当前的goroutine等待信号量
	//quit := make(chan os.Signal)
	//// 监控信号：SIcoreT, SIGTERM, SIGQUIT
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//// 这里会阻塞当前goroutine等待信号
	//<-quit
	//
	//// 调用Server.Shutdown graceful结束
	//timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//
	//if err := server.Shutdown(timeoutCtx); err != nil {
	//	log.Fatal("Server Shutdown:", err)
	//}
}
