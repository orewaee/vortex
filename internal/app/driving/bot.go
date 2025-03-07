package driving

import "context"

type Bot interface {
	Run() error
	Shutdown(ctx context.Context) error
}
