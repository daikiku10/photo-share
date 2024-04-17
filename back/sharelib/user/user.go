package user

import (
	"encoding/base64"
	"encoding/json"
	"photo-share/back/sharelib/domain/logging"
)

// ID ユーザーID
type ID string

// go:generate accessor -type=User
type User struct {
	id ID
}

func New(id ID) *User {
	return &User{
		id: id,
	}
}

// NewWithBase64Json Base64にエンコードされたユーザー情報をデコードする
func NewWithBase64Json(encoded string) (*User, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		logging.Error("ユーザー情報のBase64 Decodeに失敗しました。")
		return nil, err
	}

	// jsonからUser構造体への変換
	user, err := unmarshalJSON(decoded)
	if err != nil {
		logging.Error("ユーザー情報のJson Parseに失敗しました。")
		return nil, err
	}

	return user, nil
}

type userInformation struct {
	Id ID `json:"Id"`
}

// unmarshalJSON JSONをUser構造体に変換する
func unmarshalJSON(data []byte) (*User, error) {
	var userInfo userInformation
	if err := json.Unmarshal(data, &userInfo); err != nil {
		return nil, err
	}
	user := New(
		userInfo.Id,
	)
	return user, nil
}
