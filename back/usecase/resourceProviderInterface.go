package usecase

import (
	Repolib "photo-share/back/sharelib/domain/repository"
)

// ResourceProviderの設計書
type ResourceProviderInterface interface {
	Repolib.TransactorProvider
	// リポジトリ
	// PhotosRepository　写真リポジトリ
	PhotosRepository(trns Repolib.Transactor) (PhotoShareRepositoryInterface, error)
}
