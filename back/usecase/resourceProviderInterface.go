package usecase

// ResourceProviderの設計書
type ResourceProviderInterface interface {
	// リポジトリ
	PhotoShareRepository() error
}
