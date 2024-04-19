package usecase

import (
	repolib "photo-share/back/sharelib/domain/repository"
)

// ResourceProviderの設計書
type ResourceProviderInterface interface {
	repolib.TransactorProvider
	// リポジトリ
	PhotoShareRepository() (PhotoShareRepositoryInterface, error)
}
