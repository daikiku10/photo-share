package secret

type SecretInterface interface {
	User() string
	Password() string

	// Host ホスト名の取得
	// {hostname}:{port}形式で取得する
	Host() string
}
