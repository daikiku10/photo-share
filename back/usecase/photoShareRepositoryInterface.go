package usecase

import (
	"photo-share/back/domain/photo"
	"photo-share/back/sharelib/user"
)

type PhotoShareRepositoryInterface interface {
	// Save 写真を保存する
	Save(photo *photo.Photo, user *user.User) error
}
