package usecase

import (
	"photo-share/back/domain"
	"photo-share/back/sharelib/user"
)

type PhotoShareRepositoryInterface interface {
	// Save 写真を保存する
	Save(photo *domain.Photo, user *user.User) error
}
