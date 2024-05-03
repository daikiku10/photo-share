package photos

import (
	"photo-share/back/domain/photo"
	"photo-share/back/errorcode"
	"photo-share/back/repository/internal"
	DomainError "photo-share/back/sharelib/domain/error"
	"photo-share/back/sharelib/domain/logging"
	Repolib "photo-share/back/sharelib/domain/repository"
	"photo-share/back/sharelib/user"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type PhotosRepository struct {
	trns Repolib.Transactor
}

func NewPhotosRepository(trns Repolib.Transactor) (*PhotosRepository, error) {
	repo := &PhotosRepository{
		trns: trns,
	}
	return repo, nil
}

// Save 写真を保存する
func (pr *PhotosRepository) Save(photo *photo.Photo, user *user.User) error {
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

// FindById 指定したIDの投稿写真データを取得する
func (pr *PhotosRepository) FindById(id photo.Id, lock bool) (*photo.Photo, error) {
	mods := []qm.QueryMod{
		internal.PhotosWhere.ID.EQ(string(id)),
	}

	if lock {
		mods = append(mods, qm.For("update"))
	}

	query := internal.PluralPhotos(mods...)
	// レコード取得
	items, queryErr := query.All(pr.trns.Context(), pr.trns.Get())
	if queryErr != nil {
		return nil, DomainError.NewErrorWithInner(errorcode.SearchDBinFindById, queryErr)
	}

	if len(items) == 0 {
		return nil, nil
	}
	return ToDomain(items[0])
}

// Delete 引数で渡した投稿写真データを削除する
func (pr *PhotosRepository) Delete(photo *photo.Photo, user *user.User) error {
	// 削除対象の存在確認
	mods := []qm.QueryMod{
		internal.PhotosWhere.ID.EQ(string(photo.Id())),
		qm.For("update"),
	}

	query := internal.PluralPhotos(mods...)
	// レコード取得
	items, queryErr := query.All(pr.trns.Context(), pr.trns.Get())
	if queryErr != nil {
		return DomainError.NewErrorWithInner(errorcode.RecordDeleteFailed, queryErr)
	}
	if len(items) == 0 {
		logging.Error("削除対象が存在しません。", logging.Var("photoId", photo.Id()))
		return DomainError.NewErrorWithInner(errorcode.RecordDeleteFailed, queryErr)
	}

	// 物理削除
	internalPhoto := ToDTO(photo)
	_, deleteErr := internalPhoto.Delete(pr.trns.Context(), pr.trns.Get())
	if deleteErr != nil {
		logging.Error("削除に失敗しました。", logging.Var("photoId", photo.Id()), logging.Var("error", deleteErr))
		return DomainError.NewErrorWithInner(errorcode.RecordDeleteFailed, deleteErr)
	}

	return nil
}
