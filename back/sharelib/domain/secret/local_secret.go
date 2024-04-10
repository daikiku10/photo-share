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
