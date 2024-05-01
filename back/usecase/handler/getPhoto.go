package handler

import (
	"photo-share/back/infrastructure/restexe/apidef/server"

	"github.com/gin-gonic/gin"
)

// GetPhoto 投稿写真情報の1件取得API
// (GET /photos/{photoId})
func (s *Server) GetPhoto(c *gin.Context, photoId string, params server.GetPhotoParams) {

}
