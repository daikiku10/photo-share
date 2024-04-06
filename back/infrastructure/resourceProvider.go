package infrastructure

import "fmt"

type ResourceProvider struct {
	config Configs
}

func NewResourceProvider(config Configs) *ResourceProvider {
	return &ResourceProvider{
		config: config,
	}
}

func (rp *ResourceProvider) PhotoShareRepository() error {
	fmt.Println("テストだよ！")
	return nil
}
