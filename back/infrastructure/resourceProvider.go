package infrastructure

import (
	"context"
	"database/sql"
	"photo-share/back/repository/photos"
	Repolib "photo-share/back/sharelib/domain/repository"
	"photo-share/back/usecase"
)

type ResourceProvider struct {
	Repolib.TransactorProvider
	config Configs
}

func NewResourceProvider(
	config Configs,
	mysql *sql.DB,
	ctx context.Context,
) *ResourceProvider {
	return &ResourceProvider{
		TransactorProvider: Repolib.NewTransactorProvider(mysql, ctx),
		config:             config,
	}
}

func (rp *ResourceProvider) NewRepository(trns Repolib.Transactor) (usecase.PhotoShareRepositoryInterface, error) {
	return photos.NewPhotoShareRepository(trns)
}
