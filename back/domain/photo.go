package domain

import (
	"photo-share/back/errorcode"
	DomainError "photo-share/back/sharelib/domain/error"
	"photo-share/back/sharelib/domain/logging"
	"unicode/utf8"

	"github.com/oklog/ulid/v2"
)

type Id string

const (
	// maxTitleLength タイトルの最大文字数
	maxTitleLength = 100
)

// go:generate accessor -type=Photo
type Photo struct {
	id    Id
	title string
}

// New IDを新規採番してオブジェクトを生成する
func New(
	title string,
) (*Photo, error) {
	return NewForRepository(
		Id(ulid.Make().String()),
		title,
	)
}

// NewForRepository 全データを指定してオブジェクトの生成する
func NewForRepository(id Id, title string) (*Photo, error) {
	photo := &Photo{
		id: id,
	}

	err := photo.Edit(title)
	if err != nil {
		return nil, err
	}

	return photo, nil
}

// Edit 編集、バリデーションを行う
func (photo *Photo) Edit(title string) error {
	// バリデーションを行う
	if title == "" || utf8.RuneCountInString(title) > maxTitleLength {
		logging.Error("titleの値が空文字か101文字以上です", logging.Var("photoId", photo.id))
		return DomainError.NewError(errorcode.Validation)
	}
	photo.title = title
	return nil
}
