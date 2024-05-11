package token

import "time"

type Maker interface {
	// CreateToken 生成Token
	CreateToken(content []byte, expireDate time.Duration) (string, *Payload, error)
	// VerifyToken 解析token
	VerifyToken(token string) (*Payload, error)
}
