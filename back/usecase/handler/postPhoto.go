package handler

import (
	"fmt"
	"net/http"
	"photo-share/back/infrastructure/restexe/apidef/server"
	"photo-share/back/usecase/converter"

	"github.com/gin-gonic/gin"
)

// PostPhoto 投稿写真登録API
// (POST /photo/create)
func (s *Server) PostPhoto(ctx *gin.Context) {
	// TODO: ヘッダーにユーザー情報を受け取れるようにする
	user := ""
	// user, errHeader := user.NewWithBase64Json("")
	// if errHeader != nil {
	// 	ctx.JSON(http.StatusInternalServerError, converter.ToErrorItem(DomainError.NewErrorWithInner(ErrorCode.NotAuthorized, errHeader)))
	// 	return
	// }

	// トランザクションの取得
	transactor := s.resourceProvider.GetCurrentTransactor(ctx)
	// リポジトリの取得
	repo, errRepo := s.resourceProvider.NewRepository(transactor)
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
	fmt.Println(user, repo)
	fmt.Print(requestBody)
}
