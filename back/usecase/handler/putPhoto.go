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

// PutPhoto 投稿写真編集API
// (PUT /photos/{photoId})
func (s *Server) PutPhoto(ctx *gin.Context, photoId string) {
	// TODO: ヘッダーにユーザー情報を受け取れるようにする
	user := user.New("user1")
	// user, errHeader := user.NewWithBase64Json("")
	// if errHeader != nil {
	// 	ctx.JSON(http.StatusInternalServerError, converter.ToErrorItem(DomainError.NewErrorWithInner(ErrorCode.NotAuthorized, errHeader)))
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

	// requestBodyの取得
	var requestBody server.PutPhotoRequest
	errBind := ctx.ShouldBindJSON(&requestBody)
	if errBind != nil {
		ctx.JSON(http.StatusInternalServerError, converter.ToErrorItem(errBind))
		return
	}

	// 現在の投稿写真データの取得
	photoObject, queryErr := repo.FindById(photo.Id(photoId), false)
	if queryErr != nil {
		ctx.JSON(http.StatusInternalServerError, converter.ToErrorItem(queryErr))
		return
	}
	// 取得できなかった場合も処理終了する
	if photoObject == nil {
		ctx.JSON(http.StatusInternalServerError, converter.ToErrorItem(DomainError.NewError(errorcode.NotFound)))
		return
	}

	// 投稿写真の更新
	errUpdate := photoObject.Edit(
		requestBody.Title,
		requestBody.Description,
		requestBody.ImageUrl,
		requestBody.AuthorId,
		requestBody.CategoryId,
	)
	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, converter.ToErrorItem(errUpdate))
		return
	}

	// 投稿写真を保存
	errSave := repo.Save(photoObject, user)
	if errSave != nil {
		ctx.JSON(http.StatusInternalServerError, converter.ToErrorItem(errSave))
		return
	}

	// レスポンスの作成
	dto := server.PutPhotoSuccessResponse{
		Id: string(photoObject.Id()),
	}
	ctx.JSON(http.StatusOK, dto)

}
