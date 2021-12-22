package jwt

import "time"

type TokenData struct {
	Token    string
	ExpireAt int64
}
type AuthData struct {
	TokenInfo TokenData `json:"token_info"`
}

func Login() (error, interface{}) {
	// 业务代码。校验用户，登陆等等
	//
	//
	//
	memberID := int64(1)
	expireAt := time.Now().Unix() + 86400
	token, err := GenToken(memberID, expireAt)
	if err != nil {
		return nil, err
	}

	// 6. 返回结果
	return nil, &AuthData{
		TokenInfo: TokenData{
			Token:    token,
			ExpireAt: expireAt,
		},
	}
}
