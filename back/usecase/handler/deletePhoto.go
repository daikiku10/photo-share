package handler

import (
	"photo-share/back/infrastructure/restexe/apidef/server"

	"github.com/gin-gonic/gin"
)

// DeletePhoto 投稿写真の1件削除API
// (DELETE /photos/{photoId})
func (s *Server) DeletePhoto(c *gin.Context, photoId string, params server.DeletePhotoParams) {

}
