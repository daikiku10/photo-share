package error

import "fmt"

type ErrorCode interface {
	// AppCode アプリケーションのエラーコードを返す
	AppCode() int
	// Code エラーコードを数値変換して返す
	Code() int
}

type DomainError struct {
	Inner error
	Code  ErrorCode
}

// NewError エラーコードを元にエラーを返す
func NewError(code ErrorCode) *DomainError {
	return &DomainError{
		Inner: nil,
		Code:  code,
	}
}

// NewErrorWithInner システムエラーによるエラー生成の場合は、エラーコードを元にしたエラーに加えて、生成されたエラーを返す
func NewErrorWithInner(code ErrorCode, inner error) *DomainError {
	return &DomainError{
		Inner: inner,
		Code:  code,
	}
}

// Error エラーコードを成形して文字列で返す、システムエラーによるエラー生成の場合はその内容も返す
func (do *DomainError) Error() string {
	errMessage := "Error Code: " + do.CodeString()

	if do.Inner != nil {
		errMessage += fmt.Sprintf(", Messages: %s", do.Inner.Error())
	}
	return errMessage
}

// CodeString エラーコードを成形して文字列で返す
func (do *DomainError) CodeString() string {
	return fmt.Sprintf("%03d-%06d", do.Code.AppCode(), do.Code.Code())
}
