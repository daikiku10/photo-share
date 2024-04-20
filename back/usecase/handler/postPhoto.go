package handler

import (
	"net/http"
	ErrorCode "photo-share/back/errorcode"
	DomainError "photo-share/back/sharelib/domain/error"
	"photo-share/back/sharelib/user"
	"photo-share/back/usecase/converter"

	"github.com/gin-gonic/gin"
)

// PostPhoto 投稿写真登録API
// (POST /photo/create)
func (s *Server) PostPhoto(ctx *gin.Context) {
	// TODO: ヘッダーにユーザー情報を受け取れるようにする
	user, err := user.NewWithBase64Json("")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, converter.ToErrorItem(DomainError.NewErrorWithInner(ErrorCode.NotAuthorized, err)))
		return
	}

	// トランザクションの取得
	transactor := s.resourceProvider.GetCurrentTransactor(ctx)
	// リポジトリの取得
	repo, err := s.resourceProvider.NewRepository(transactor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, converter.ToErrorItem(err))
		return
	}

	// requestBodyを取得

}
