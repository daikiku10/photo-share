package usecase

import (
	"photo-share/back/domain/photo"
	"photo-share/back/sharelib/user"
)

type PhotoShareRepositoryInterface interface {
	// Save 写真を保存する
	Save(photo *photo.Photo, user *user.User) error
	// FindById 指定したIDの投稿写真データを取得する
	FindById(id photo.Id, lock bool) (*photo.Photo, error)
}
