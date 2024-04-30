package photos

import (
	"photo-share/back/domain/photo"
	"photo-share/back/repository/internal"
)

// ToDomain ドメインへ変換する
func ToDomain(src *internal.Photos) (*photo.Photo, error) {
	return photo.NewForRepository(
		photo.Id(src.ID),
		src.Title,
		src.Description,
		src.ImageUrl,
		src.AuthorId,
		src.CategoryId,
	)
}

// ToDTO DTOに変換する
func ToDTO(src *photo.Photo) *internal.Photos {
	return &internal.Photos{
		ID:          string(src.Id()),
		Title:       src.Title(),
		Description: src.Description(),
		ImageUrl:    src.ImageUrl(),
		AuthorId:    src.AuthorId(),
		CategoryId:  src.CategoryId(),
	}
}
