package gpxrotate_test

import (
	"context"
	"testing"
	"time"

	"github.com/glynternet/gpx-rotate/pkg/gpxrotate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContextWithTimeoutHandler(t *testing.T) {
	t.Run("should be done and return cancellation when parent cancelled", func(t *testing.T) {
		parent, cancelParent := context.WithCancel(context.Background())
		withTimeoutHandler, cancelWithTimeoutHandler := gpxrotate.WithTimeoutHandler(parent, time.Minute, func() {
			require.FailNow(t, "handler should not have been called")
		})
		defer cancelWithTimeoutHandler()
		cancelParent()
		<-withTimeoutHandler.Done()
		assert.EqualError(t, withTimeoutHandler.Err(), "context canceled")
	})

	t.Run("should be done and return cancellation when self cancelled", func(t *testing.T) {
		parent, cancelParent := context.WithCancel(context.Background())
		defer cancelParent()
		withTimeoutHandler, cancelWithTimeoutHandler := gpxrotate.WithTimeoutHandler(parent, time.Minute, func() {
			require.FailNow(t, "handler should not have been called")
		})
		cancelWithTimeoutHandler()
		<-withTimeoutHandler.Done()
		assert.EqualError(t, withTimeoutHandler.Err(), "context canceled")
	})

	t.Run("should defer value to parent context", func(t *testing.T) {
		ctx, cancel := gpxrotate.WithTimeoutHandler(
			context.WithValue(context.Background(), "foo", "bar"), time.Minute, func() {
				require.FailNow(t, "handler should not have been called")
			})
		defer cancel()
		assert.Equal(t, ctx.Value("foo").(string), "bar")
	})

	t.Run("should defer deadline to parent context if sooner than timeout deadline", func(t *testing.T) {
		deadlineIn := time.Now().Add(time.Minute)
		parent, cancelParent := context.WithDeadline(context.Background(), deadlineIn)
		defer cancelParent()
		ctx, cancel := gpxrotate.WithTimeoutHandler(parent, time.Hour, func() {
			require.FailNow(t, "handler should not have been called")
		})
		defer cancel()
		deadlineOut, ok := ctx.Deadline()
		require.True(t, ok)
		assert.Equal(t, deadlineIn, deadlineOut)
	})

	t.Run("should return timeout deadline if parent has no deadline", func(t *testing.T) {
		ctx, cancel := gpxrotate.WithTimeoutHandler(context.Background(), time.Minute, func() {
			require.FailNow(t, "handler should not have been called")
		})
		defer cancel()
		deadline, ok := ctx.Deadline()
		require.True(t, ok)
		assert.InDelta(t, 0, deadline.Sub(time.Now().Add(time.Minute)), float64(time.Second))
	})

	t.Run("should return timeout deadline if sooner than parent deadline", func(t *testing.T) {
		parent, cancelParent := context.WithTimeout(context.Background(), time.Hour)
		defer cancelParent()
		ctx, cancel := gpxrotate.WithTimeoutHandler(parent, time.Second, func() {
			require.FailNow(t, "handler should not have been called")
		})
		defer cancel()
		deadline, ok := ctx.Deadline()
		require.True(t, ok)
		assert.InDelta(t, 0, deadline.Sub(time.Now().Add(time.Second)), float64(50*time.Millisecond))
	})

	t.Run("should be done after timeout reached and return timeout error", func(t *testing.T) {
		ctx, _ := gpxrotate.WithTimeoutHandler(context.Background(), 0, func() {})
		select {
		case <-ctx.Done():
		case <-time.After(50 * time.Millisecond):
			require.FailNow(t, "should have timed out by now")
		}
		assert.EqualError(t, ctx.Err(), "timeout something")
	})

	t.Run("should call handler on timeout before closing done channel", func(t *testing.T) {
		ch := make(chan struct{})
		ctx, cancel := gpxrotate.WithTimeoutHandler(context.Background(), 10*time.Millisecond, func() {
			ch <- struct{}{}
		})
		defer cancel()
		select {
		case <-ctx.Done():
			require.FailNow(t, "handler should close other channel before context done")
		case <-ch:
		}
	})
}
