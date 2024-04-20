package errorcode

type PhotoShareErrorCode int

// AppCode アプリケーションのエラーコードを返す
func (phe PhotoShareErrorCode) AppCode() int {
	return 100
}

// Code エラーコードを数値変換して返す
func (phe PhotoShareErrorCode) Code() int {
	return int(phe)
}

// アプリケーションエラー定義
const (
	Unknown       PhotoShareErrorCode = iota + 1 // 原因不明のエラー
	NotAuthorized                                // 必要な権限がない時のエラー
)
