package app

import (
	"skyscraper/framework"
	"skyscraper/framework/containerService/contract"
)

// AppProvider 提供App的具体实现方法
type AppProvider struct {
	BaseFolder string
}

// Register 注册App方法
func (app *AppProvider) Register(container framework.Container) framework.NewInstance {
	return NewApp
}

// Boot 启动调用
func (app *AppProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (app *AppProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (app *AppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, app.BaseFolder}
}

// Name 获取字符串凭证
func (app *AppProvider) Name() string {
	return contract.AppKey
}
