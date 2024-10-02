package infrastructure

import (
	"context"
	"time"
)

type KeyValueStore interface {
	Get(context.Context, string) (string, error)
	Set(context.Context, string, string, time.Duration) error
	Health(context.Context) error
}
