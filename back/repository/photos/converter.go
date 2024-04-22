package photos

import (
	"photo-share/back/domain"
	"photo-share/back/repository/internal"
)

// ToDTO DTOに変換する
func ToDTO(src *domain.Photo) *internal.Photos {
	return &internal.Photos{
		ID:    string(src.Id()),
		Title: src.Title(),
	}
}
