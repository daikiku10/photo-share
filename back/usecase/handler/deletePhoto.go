package handler

import (
	"net/http"
	"photo-share/back/domain/photo"
	"photo-share/back/errorcode"
	"photo-share/back/infrastructure/restexe/apidef/server"
	DomainError "photo-share/back/sharelib/domain/error"
	"photo-share/back/sharelib/user"
	"photo-share/back/usecase/converter"

	"github.com/gin-gonic/gin"
)

// DeletePhoto 投稿写真の1件削除API
// (DELETE /photos/{photoId})
func (s *Server) DeletePhoto(ctx *gin.Context, photoId string, params server.DeletePhotoParams) {
	user := user.New("user1")
	// user, errHeader := user.NewWithBase64Json(params.UserInformation)
	// if errHeader != nil {
	// 	ctx.JSON(http.StatusInternalServerError, converter.ToErrorItem(DomainError.NewErrorWithInner(errorcode.NotAuthorized, errHeader)))
	// 	return
	// }

	// トランザクションの取得
	transactor := s.resourceProvider.GetCurrentTransactor(ctx)
	// リポジトリの取得
	repo, errRepo := s.resourceProvider.PhotosRepository(transactor)
	if errRepo != nil {
		ctx.JSON(http.StatusInternalServerError, converter.ToErrorItem(errRepo))
		return
	}

	// 削除対象の投稿写真データの取得
	photoObject, queryErr := repo.FindById(photo.Id(photoId), true)
	if queryErr != nil {
		ctx.JSON(http.StatusInternalServerError, converter.ToErrorItem(queryErr))
		return
	}
	// 取得できなかった場合も処理終了する
	if photoObject == nil {
		ctx.JSON(http.StatusInternalServerError, converter.ToErrorItem(DomainError.NewError(errorcode.NotFound)))
		return
	}

	// 削除処理
	deleteErr := repo.Delete(photoObject, user)
	if deleteErr != nil {
		ctx.JSON(http.StatusInternalServerError, converter.ToErrorItem(deleteErr))
		return
	}

	ctx.JSON(http.StatusOK, "削除が正常に終了しました。")
}
