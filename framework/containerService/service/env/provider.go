package env

import (
	"github.com/SKYBroGardenLush/skyscraper/framework"
	"github.com/SKYBroGardenLush/skyscraper/framework/containerService/contract"
)

type EnvProvider struct {
	Folder string
}

// Register a new function for make a service instance
func (provider *EnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewEnv
}

// Boot will be called when the service instantiate
func (provider *EnvProvider) Boot(c framework.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	provider.Folder = app.BaseFolder()
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *EnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *EnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.Folder}
}

// Name define the name for this service
func (provider *EnvProvider) Name() string {
	return contract.EnvKey
}
