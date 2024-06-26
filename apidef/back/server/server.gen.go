// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// DomainError defines model for DomainError.
type DomainError struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// GetPhotoSuccessResponse defines model for GetPhotoSuccessResponse.
type GetPhotoSuccessResponse struct {
	// AuthorId 投稿者ID
	AuthorId string `json:"authorId"`

	// CategoryId カテゴリID
	CategoryId string `json:"categoryId"`

	// CreatedAt 作成日時
	CreatedAt time.Time `json:"createdAt"`

	// Description 写真の説明
	Description string `json:"description"`

	// Id 投稿写真ID
	Id string `json:"id"`

	// ImageUrl 写真URL
	ImageUrl string `json:"imageUrl"`

	// LikedCount いいね数
	LikedCount int `json:"likedCount"`

	// Title 写真タイトル
	Title string `json:"title"`
}

// PostPhotoRequest defines model for PostPhotoRequest.
type PostPhotoRequest struct {
	// AuthorId 投稿者ID
	AuthorId string `json:"authorId"`

	// CategoryId カテゴリID
	CategoryId string `json:"categoryId"`

	// Description 写真の説明
	Description string `json:"description"`

	// ImageUrl 写真URL
	ImageUrl string `json:"imageUrl"`

	// Title 写真タイトル
	Title string `json:"title"`
}

// PostPhotoSuccessResponse defines model for PostPhotoSuccessResponse.
type PostPhotoSuccessResponse struct {
	// Id 登録した写真ID
	Id string `json:"id"`
}

// PutPhotoRequest defines model for PutPhotoRequest.
type PutPhotoRequest struct {
	// AuthorId 投稿者ID
	AuthorId string `json:"authorId"`

	// CategoryId カテゴリID
	CategoryId string `json:"categoryId"`

	// Description 写真の説明
	Description string `json:"description"`

	// Id 投稿写真ID
	Id string `json:"id"`

	// ImageUrl 写真URL
	ImageUrl string `json:"imageUrl"`

	// Title 写真タイトル
	Title string `json:"title"`
}

// PutPhotoSuccessResponse defines model for PutPhotoSuccessResponse.
type PutPhotoSuccessResponse struct {
	// Id 編集した写真ID
	Id string `json:"id"`
}

// UserInformationHeader defines model for UserInformationHeader.
type UserInformationHeader = string

// PostPhotoParams defines parameters for PostPhoto.
type PostPhotoParams struct {
	// UserInformation ユーザー情報
	UserInformation *UserInformationHeader `json:"userInformation,omitempty"`
}

// DeletePhotoParams defines parameters for DeletePhoto.
type DeletePhotoParams struct {
	// UserInformation ユーザー情報
	UserInformation *UserInformationHeader `json:"userInformation,omitempty"`
}

// GetPhotoParams defines parameters for GetPhoto.
type GetPhotoParams struct {
	// UserInformation ユーザー情報
	UserInformation *UserInformationHeader `json:"userInformation,omitempty"`
}

// PutPhotoParams defines parameters for PutPhoto.
type PutPhotoParams struct {
	// UserInformation ユーザー情報
	UserInformation *UserInformationHeader `json:"userInformation,omitempty"`
}

// PostPhotoJSONRequestBody defines body for PostPhoto for application/json ContentType.
type PostPhotoJSONRequestBody = PostPhotoRequest

// PutPhotoJSONRequestBody defines body for PutPhoto for application/json ContentType.
type PutPhotoJSONRequestBody = PutPhotoRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 投稿写真登録API
	// (POST /photos)
	PostPhoto(c *gin.Context, params PostPhotoParams)
	// 投稿写真の1件削除API
	// (DELETE /photos/{photoId})
	DeletePhoto(c *gin.Context, photoId string, params DeletePhotoParams)
	// 投稿写真情報の1件取得API
	// (GET /photos/{photoId})
	GetPhoto(c *gin.Context, photoId string, params GetPhotoParams)
	// 投稿写真編集API
	// (PUT /photos/{photoId})
	PutPhoto(c *gin.Context, photoId string, params PutPhotoParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// PostPhoto operation middleware
func (siw *ServerInterfaceWrapper) PostPhoto(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params PostPhotoParams

	headers := c.Request.Header

	// ------------- Optional header parameter "userInformation" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("userInformation")]; found {
		var UserInformation UserInformationHeader
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandler(c, fmt.Errorf("Expected one value for userInformation, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "userInformation", runtime.ParamLocationHeader, valueList[0], &UserInformation)
		if err != nil {
			siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter userInformation: %w", err), http.StatusBadRequest)
			return
		}

		params.UserInformation = &UserInformation

	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostPhoto(c, params)
}

// DeletePhoto operation middleware
func (siw *ServerInterfaceWrapper) DeletePhoto(c *gin.Context) {

	var err error

	// ------------- Path parameter "photoId" -------------
	var photoId string

	err = runtime.BindStyledParameter("simple", false, "photoId", c.Param("photoId"), &photoId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter photoId: %w", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params DeletePhotoParams

	headers := c.Request.Header

	// ------------- Optional header parameter "userInformation" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("userInformation")]; found {
		var UserInformation UserInformationHeader
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandler(c, fmt.Errorf("Expected one value for userInformation, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "userInformation", runtime.ParamLocationHeader, valueList[0], &UserInformation)
		if err != nil {
			siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter userInformation: %w", err), http.StatusBadRequest)
			return
		}

		params.UserInformation = &UserInformation

	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeletePhoto(c, photoId, params)
}

// GetPhoto operation middleware
func (siw *ServerInterfaceWrapper) GetPhoto(c *gin.Context) {

	var err error

	// ------------- Path parameter "photoId" -------------
	var photoId string

	err = runtime.BindStyledParameter("simple", false, "photoId", c.Param("photoId"), &photoId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter photoId: %w", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPhotoParams

	headers := c.Request.Header

	// ------------- Optional header parameter "userInformation" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("userInformation")]; found {
		var UserInformation UserInformationHeader
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandler(c, fmt.Errorf("Expected one value for userInformation, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "userInformation", runtime.ParamLocationHeader, valueList[0], &UserInformation)
		if err != nil {
			siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter userInformation: %w", err), http.StatusBadRequest)
			return
		}

		params.UserInformation = &UserInformation

	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetPhoto(c, photoId, params)
}

// PutPhoto operation middleware
func (siw *ServerInterfaceWrapper) PutPhoto(c *gin.Context) {

	var err error

	// ------------- Path parameter "photoId" -------------
	var photoId string

	err = runtime.BindStyledParameter("simple", false, "photoId", c.Param("photoId"), &photoId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter photoId: %w", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params PutPhotoParams

	headers := c.Request.Header

	// ------------- Optional header parameter "userInformation" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("userInformation")]; found {
		var UserInformation UserInformationHeader
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandler(c, fmt.Errorf("Expected one value for userInformation, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "userInformation", runtime.ParamLocationHeader, valueList[0], &UserInformation)
		if err != nil {
			siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter userInformation: %w", err), http.StatusBadRequest)
			return
		}

		params.UserInformation = &UserInformation

	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutPhoto(c, photoId, params)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/photos", wrapper.PostPhoto)
	router.DELETE(options.BaseURL+"/photos/:photoId", wrapper.DeletePhoto)
	router.GET(options.BaseURL+"/photos/:photoId", wrapper.GetPhoto)
	router.PUT(options.BaseURL+"/photos/:photoId", wrapper.PutPhoto)
}
