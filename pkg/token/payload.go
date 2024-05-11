package token

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	// 用于管理每个JWT
	ID      uuid.UUID
	Content []byte
	// 创建时间用于检验
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}
