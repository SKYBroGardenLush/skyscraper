package kernel

import (
	"net/http"
	"github.com/SKYBroGardenLush/skyscraper/framework"
)

// KernelService 引擎服务
type KernelService struct {
	core *framework.Core
}

// NewKernelService 初始化web引擎服务实例
func NewKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*framework.Core)
	return &KernelService{core: httpEngine}, nil
}

// HttpEngine 返回web引擎
func (s *KernelService) HttpEngine() http.Handler {
	return s.core
}
