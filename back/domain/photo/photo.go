package photo

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
	// maxDescriptionLength 説明のの最大文字数
	maxDescriptionLength = 100
)

// go:generate accessor -type=Photo
type Photo struct {
	id          Id
	title       string
	description string
	imageUrl    string
	authorId    string
	categoryId  string
}

// New IDを新規採番してオブジェクトを生成する
func New(
	title string,
	description string,
	imageUrl string,
	authorId string,
	categoryId string,
) (*Photo, error) {
	return NewForRepository(
		Id(ulid.Make().String()),
		title,
		description,
		imageUrl,
		authorId,
		categoryId,
	)
}

// NewForRepository 全データを指定してオブジェクトの生成する
func NewForRepository(
	id Id,
	title string,
	description string,
	imageUrl string,
	authorId string,
	categoryId string,
) (*Photo, error) {
	photo := &Photo{
		id: id,
	}

	err := photo.Edit(
		title,
		description,
		imageUrl,
		authorId,
		categoryId,
	)
	if err != nil {
		return nil, err
	}

	return photo, nil
}

// Edit 編集、バリデーションを行う
func (photo *Photo) Edit(
	title string,
	description string,
	imageUrl string,
	authorId string,
	categoryId string,
) error {
	// バリデーションを行う
	if title == "" || utf8.RuneCountInString(title) > maxTitleLength {
		logging.Error("titleの値が空文字か101文字以上です", logging.Var("photoId", photo.id))
		return DomainError.NewError(errorcode.Validation)
	}
	if utf8.RuneCountInString(description) > maxDescriptionLength {
		logging.Error("descriptionの値が空文字か101文字以上です", logging.Var("photoId", photo.id))
		return DomainError.NewError(errorcode.Validation)
	}
	photo.title = title
	photo.description = description
	photo.imageUrl = imageUrl
	photo.authorId = authorId
	photo.categoryId = categoryId
	return nil
}
