package infrastructure

import (
	"photo-share/back/repository"
	"photo-share/back/usecase"
)

type ResourceProvider struct {
	config Configs
}

func NewResourceProvider(config Configs) *ResourceProvider {
	return &ResourceProvider{
		config: config,
	}
}

func (rp *ResourceProvider) PhotoShareRepository() (usecase.PhotoShareRepositoryInterface, error) {
	return repository.NewPhotoShareRepository()
}
