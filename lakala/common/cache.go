package common

import (
	"context"
	"time"
)

type ICache interface {
	GetAccessToken(ctx context.Context, key string) (string, error)
	SetAccessToken(ctx context.Context, key, token string, expireTime time.Duration) error
	DelAccessToken(ctx context.Context, key string) error
}
