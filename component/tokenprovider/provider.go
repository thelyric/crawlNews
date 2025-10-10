package tokenprovider

import (
	"my-app/common"
	"time"
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

var (
	ErrNotFound      = common.MakeFailResponse("ERR_NOT_FOUND", "Token not found")
	ErrEncodingToken = common.MakeFailResponse("ERR_ENDCOING_TOKEN", "Error encoding the token")
	ErrInvalidToken  = common.MakeFailResponse("ERR_INVALID_TOKEN", "Invalid token provided")
)
