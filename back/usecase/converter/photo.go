package converter

import (
	"photo-share/back/domain"
	"photo-share/back/infrastructure/restexe/apidef/server"
	"photo-share/back/sharelib/user"
)

// ToDomain リクエスト用のDTOをドメインオブジェクトに変換する
func ToDomain(requestBody server.PostPhotoRequest, user *user.User) (*domain.Photo, error) {
	return domain.New(requestBody.Title)
}
