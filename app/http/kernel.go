package http

import "skyscraper/framework"

// NewHttpCore 创建了一个绑定了路由的Web引擎
func NewHttpCore() (*framework.Core, error) {

	return framework.NewCore(), nil
}