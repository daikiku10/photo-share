package error

type ErrorCode interface {
	AppCode() int
	Code() int
}

type DomainError struct {
	Inner error
	Code  ErrorCode
}

// NewErrorWithInner システムエラーによるエラー生成の場合は、エラーコードを元にしたエラーに加えて、生成されたエラーを返す
func NewErrorWithInner(code ErrorCode, inner error) *DomainError {
	return &DomainError{
		Inner: inner,
		Code:  code,
	}
}
