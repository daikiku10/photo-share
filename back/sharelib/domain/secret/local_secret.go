package secret

type LocalSecret struct {
	user     string
	password string
	host     string
}

// NewLocalSecret ローカルのシークレット情報構造体の作成
func NewLocalSecret(user, password, host string) *LocalSecret {
	return &LocalSecret{
		user:     user,
		password: password,
		host:     host,
	}
}

// User ユーザー名の取得
func (ls *LocalSecret) User() string {
	return ls.user
}

// Password パスワードの取得
func (ls *LocalSecret) Password() string {
	return ls.password
}

// Host ホスト名の取得
func (ls *LocalSecret) Host() string {
	return ls.host
}
