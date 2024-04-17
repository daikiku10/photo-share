package handler

import (
	"fmt"
	"net/http"
	"photo-share/back/sharelib/user"

	"github.com/gin-gonic/gin"
)

// PostPhoto 投稿写真登録API
// (POST /photo/create)
func (s *Server) PostPhoto(ctx *gin.Context) {
	// TODO: ヘッダーにユーザー情報を受け取れるようにする
	user, err := user.NewWithBase64Json("")
	if err != nil {
		// TODO: エラー内容の返却
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	fmt.Println(user)
}
