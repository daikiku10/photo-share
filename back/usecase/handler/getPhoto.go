package handler

import (
	"net/http"
	"photo-share/back/domain/photo"
	"photo-share/back/errorcode"
	"photo-share/back/infrastructure/restexe/apidef/server"
	DomainError "photo-share/back/sharelib/domain/error"
	"photo-share/back/usecase/converter"

	"github.com/gin-gonic/gin"
)

// GetPhoto 投稿写真情報の1件取得API
// (GET /photos/{photoId})
func (s *Server) GetPhoto(ctx *gin.Context, photoId string, params server.GetPhotoParams) {
	// トランザクションの取得
	transactor := s.resourceProvider.GetCurrentTransactor(ctx)
	// リポジトリの取得
	repo, errRepo := s.resourceProvider.PhotosRepository(transactor)
	if errRepo != nil {
		ctx.JSON(http.StatusInternalServerError, converter.ToErrorItem(errRepo))
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

	// レスポンスの作成
	dto := server.GetPhotoSuccessResponse{
		Id:          string(photoObject.Id()),
		Title:       string(photoObject.Title()),
		Description: string(photoObject.Description()),
		ImageUrl:    string(photoObject.ImageUrl()),
		AuthorId:    string(photoObject.AuthorId()),
		CategoryId:  string(photoObject.CategoryId()),
		CreatedAt:   photoObject.CreatedAt(),
		LikedCount:  0, // TODO: 固定値ではない正しい値を返却す
	}

	ctx.JSON(http.StatusOK, dto)
}
