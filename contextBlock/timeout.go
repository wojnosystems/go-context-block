package contextBlock

import (
	"context"
	"time"
)

// WithTimeout creates a new context with the timeout added
func WithTimeout(ctx context.Context, timeout time.Duration, callback func(ctx context.Context)) {
	blockCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	callback(blockCtx)
}
