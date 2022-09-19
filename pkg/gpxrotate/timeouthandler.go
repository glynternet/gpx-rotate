package gpxrotate

import (
	"context"
	"errors"
	"sync"
	"time"
)

// WithTimeoutHandler creates a context that will call the given handler function, handle, after the timeout
// duration has exceeded. The handle function will be called and the context will wait for it to return before
// closing the done channel.
// If the returned context is cancelled during handling of the timeout handle, the Done channel will be
// closed and the error returned will be a context cancellation error.
func WithTimeoutHandler(ctx context.Context, timeout time.Duration, handle func()) (context.Context, context.CancelFunc) {
	handler := contextWithTimeoutHandler{
		parent:   ctx,
		deadline: time.Now().Add(timeout),
		doneCh:   make(chan struct{}),
	}

	go func() {
		select {
		case <-ctx.Done():
			handler.close(2)
		case <-time.After(timeout):
			handle()
			handler.close(3)
		}
	}()

	return &handler, func() { handler.close(1) }
}

type contextWithTimeoutHandler struct {
	parent      context.Context
	deadline    time.Time
	closeOnce   sync.Once
	doneCh      chan struct{}
	closeReason uint8
}

func (ctx *contextWithTimeoutHandler) close(reason uint8) {
	ctx.closeOnce.Do(func() {
		ctx.closeReason = reason
		close(ctx.doneCh)
	})
}

func (c *contextWithTimeoutHandler) Deadline() (time.Time, bool) {
	parentDeadline, ok := c.parent.Deadline()
	if !ok {
		return c.deadline, true
	}
	if c.deadline.Before(parentDeadline) {
		return c.deadline, ok
	}
	return parentDeadline, ok
}

func (c *contextWithTimeoutHandler) Done() <-chan struct{} {
	return c.doneCh
}

func (c *contextWithTimeoutHandler) Err() error {
	switch c.closeReason {
	case 1:
		return errors.New("context canceled")
	case 2:
		return c.parent.Err()
	case 3:
		return errors.New("timeout something")
	}
	return nil
}

func (c *contextWithTimeoutHandler) Value(key interface{}) interface{} {
	return c.parent.Value(key)
}
