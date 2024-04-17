package user

// ID ユーザーID
type ID string

// go:generate accessor -type=User
type User struct {
	id ID
}
