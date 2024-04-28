package converter

import (
	"photo-share/back/domain/photo"
	"photo-share/back/infrastructure/restexe/apidef/server"
	"photo-share/back/sharelib/user"
)

// ToDomain リクエスト用のDTOをドメインオブジェクトに変換する
func ToDomain(requestBody server.PostPhotoRequest, user *user.User) (*photo.Photo, error) {
	return photo.New(
		requestBody.Title,
		requestBody.Description,
		requestBody.ImageUrl,
		requestBody.AuthorId,
		requestBody.CategoryId,
	)
}
