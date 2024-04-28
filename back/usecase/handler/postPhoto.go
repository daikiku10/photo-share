package handler

import (
	"net/http"
	"photo-share/back/infrastructure/restexe/apidef/server"
	"photo-share/back/sharelib/user"
	"photo-share/back/usecase/converter"

	"github.com/gin-gonic/gin"
)

// PostPhoto 投稿写真登録API
// (POST /photos)
func (s *Server) PostPhoto(ctx *gin.Context) {
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

	// requestBodyを取得
	var requestBody server.PostPhotoRequest
	errBind := ctx.ShouldBindJSON(&requestBody)
	if errBind != nil {
		ctx.JSON(http.StatusInternalServerError, converter.ToErrorItem(errBind))
		return
	}

	// ドメインに変換する
	photoObject, errNew := converter.ToDomain(requestBody, user)
	if errNew != nil {
		ctx.JSON(http.StatusInternalServerError, converter.ToErrorItem(errNew))
		return
	}

	errSave := repo.Save(photoObject, user)
	if errSave != nil {
		ctx.JSON(http.StatusInternalServerError, converter.ToErrorItem(errSave))
		return
	}

	dto := server.PostPhotoSuccessResponse{
		Id: string(photoObject.Id()),
	}
	ctx.JSON(http.StatusOK, dto)
}
