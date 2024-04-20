package converter

import (
	ErrorCode "photo-share/back/errorcode"
	"photo-share/back/infrastructure/restexe/apidef/server"
	DomainError "photo-share/back/sharelib/domain/error"
)

// ToErrorItem errorをDomainError構造に変換する
func ToErrorItem(err error) server.DomainError {
	localErr := DomainError.NewErrorWithInner(ErrorCode.Unknown, err)
	convert, ok := err.(*DomainError.DomainError)
	if ok {
		localErr = convert
	}

	code := localErr.CodeString()
	message := localErr.Error()

	return server.DomainError{
		Code:    &code,
		Message: &message,
	}
}
