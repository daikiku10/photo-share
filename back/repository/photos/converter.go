package photos

import (
	"photo-share/back/domain/photo"
	"photo-share/back/repository/internal"
)

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
