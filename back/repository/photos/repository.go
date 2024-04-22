package photos

import (
	"photo-share/back/domain"
	"photo-share/back/errorcode"
	"photo-share/back/repository/internal"
	DomainError "photo-share/back/sharelib/domain/error"
	"photo-share/back/sharelib/domain/logging"
	Repolib "photo-share/back/sharelib/domain/repository"
	"photo-share/back/sharelib/user"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type PhotoShareRepository struct {
	trns Repolib.Transactor
}

func NewPhotoShareRepository(trns Repolib.Transactor) (*PhotoShareRepository, error) {
	repo := &PhotoShareRepository{
		trns: trns,
	}
	return repo, nil
}

// Save 写真を保存する
func (pr *PhotoShareRepository) Save(photo *domain.Photo, user *user.User) error {
	mods := []qm.QueryMod{
		internal.PhotosWhere.ID.EQ(string(photo.Id())),
		qm.For("update"),
	}

	query := internal.PluralPhotos(mods...)
	// レコード取得
	items, queryErr := query.All(pr.trns.Context(), pr.trns.Get())
	if queryErr != nil {
		return DomainError.NewErrorWithInner(errorcode.RecordInsertFailed, queryErr)
	}

	// DTOに変換
	internalPhotos := ToDTO(photo)

	// Idが存在しない場合は新規作成
	if len(items) == 0 {
		err := internalPhotos.Insert(pr.trns.Context(), pr.trns.Get(), boil.Infer())
		if err != nil {
			logging.Error("photosテーブルの作成に失敗しました。", logging.Var("error", err))
			return DomainError.NewError(errorcode.RecordInsertFailed)
		}
	} else {
		_, err := internalPhotos.Update(pr.trns.Context(), pr.trns.Get(), boil.Infer())
		if err != nil {
			logging.Error("photosテーブルの更新に失敗しました。", logging.Var("error", err))
			return DomainError.NewError(errorcode.RecordUpdateFailed)
		}
	}
	return nil
}
