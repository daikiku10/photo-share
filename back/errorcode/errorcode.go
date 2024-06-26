package errorcode

type PhotoShareErrorCode int

// AppCode アプリケーションのエラーコードを返す
func (phe PhotoShareErrorCode) AppCode() int {
	return 101
}

// Code エラーコードを数値変換して返す
func (phe PhotoShareErrorCode) Code() int {
	return int(phe)
}

// アプリケーションエラー定義
const (
	Unknown            PhotoShareErrorCode = iota + 1 // 原因不明のエラー
	NotAuthorized                                     // 必要な権限がない時のエラー
	Validation                                        // バリデーションエラー
	RecordInsertFailed                                // Saveメソッド内で、データの作成に失敗した時のエラー
	RecordUpdateFailed                                // Saveメソッド内で、データの更新に失敗した時のエラー
	SearchDBinFindById                                // FindByIDメソッド内で、DB検索に失敗した時のエラー
	NotFound                                          // データが見つからないエラー
	RecordDeleteFailed                                // Deleteメソッド内で、データの削除に失敗した時のエラー                                    // データが見つからないエラー
)
