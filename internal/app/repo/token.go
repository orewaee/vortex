package repo

import (
	"context"
	"time"
)

type TokenReader interface {
	TokenExists(ctx context.Context, token string) (bool, error)
}

type TokenWriter interface {
	AddToken(ctx context.Context, token string, lifetime time.Duration) error
	RemoveToken(ctx context.Context, token string) error
}

type TokenReadWriter interface {
	TokenReader
	TokenWriter
}
