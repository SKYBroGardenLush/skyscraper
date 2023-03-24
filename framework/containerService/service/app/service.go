package app

import (
	"errors"
	"flag"
	"path/filepath"
	"github.com/SKYBroGardenLush/skyscraper/framework"
	"github.com/SKYBroGardenLush/skyscraper/framework/utils"
)

// App 代表skyscraper框架的App实现
type App struct {
	container  framework.Container // 服务容器
	baseFolder string              // 基础路径
	configMap  map[string]string   // 配置加载
}

// Version 实现版本
func (a App) Version() string {
	return "0.0.3"
}

// BaseFolder 表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (a App) BaseFolder() string {
	if a.baseFolder != "" {
		return a.baseFolder
	}

	// 如果没有设置，则使用参数
	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数, 默认为当前路径")
	flag.Parse()
	if baseFolder != "" {
		return baseFolder
	}

	// 如果参数也没有，使用默认的当前路径
	return utils.GetExecDirectory()
}

// ConfigFolder  表示配置文件地址
func (a App) ConfigFolder() string {
	return filepath.Join(a.BaseFolder(), "config")
}

// LogFolder 表示日志存放地址
func (a App) LogFolder() string {
	return filepath.Join(a.StorageFolder(), "log")
}

func (a App) HttpFolder() string {
	return filepath.Join(a.BaseFolder(), "http")
}

func (a App) ConsoleFolder() string {
	return filepath.Join(a.BaseFolder(), "console")
}

func (a App) StorageFolder() string {
	return filepath.Join(a.BaseFolder(), "storage")
}

// ProviderFolder 定义业务自己的服务提供者地址
func (a App) ProviderFolder() string {
	return filepath.Join(a.BaseFolder(), "provider")
}

// MiddlewareFolder 定义业务自己定义的中间件
func (a App) MiddlewareFolder() string {
	return filepath.Join(a.HttpFolder(), "middleware")
}

// CommandFolder 定义业务定义的命令
func (a App) CommandFolder() string {
	return filepath.Join(a.ConsoleFolder(), "command")
}

// RuntimeFolder 定义业务的运行中间态信息
func (a App) RuntimeFolder() string {
	return filepath.Join(a.StorageFolder(), "runtime")
}

// TestFolder 定义测试需要的信息
func (a App) TestFolder() string {
	return filepath.Join(a.BaseFolder(), "test")
}

// DeployFolder 定义测试需要的信息
func (a App) DeployFolder() string {
	if val, ok := a.configMap["deploy_folder"]; ok {
		return val
	}
	return filepath.Join(a.BaseFolder(), "deploy")
}

// LoadAppConfig 加载配置map
func (a App) LoadAppConfig(kv map[string]string) {
	for key, val := range kv {
		a.configMap[key] = val
	}
}

// AppFolder 代表app目录
func (a App) AppFolder() string {
	if val, ok := a.configMap["app_folder"]; ok {
		return val
	}
	return filepath.Join(a.BaseFolder(), "app")
}

// NewApp 初始化App
func NewApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	// 有两个参数，一个是容器，一个是baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return &App{baseFolder: baseFolder, container: container}, nil
}
