// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package server

import (
	"github.com/gin-gonic/gin"
)

// DomainError defines model for DomainError.
type DomainError struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
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

// PostPhotoJSONRequestBody defines body for PostPhoto for application/json ContentType.
type PostPhotoJSONRequestBody = PostPhotoRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 投稿写真登録API
	// (POST /photos)
	PostPhoto(c *gin.Context)
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

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostPhoto(c)
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
}
