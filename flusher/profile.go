package flusher

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

var (
	taskCount uint32
)

type flushOption struct {
	ctx context.Context

	TaskName        string
	BufferSize      int
	IntervalInMills time.Duration
	Concurrent      int
	QueueSize       int
}

type FlushOption func(option *flushOption)

func defaultOption() *flushOption {
	option := &flushOption{
		ctx:             context.Background(),
		TaskName:        fmt.Sprintf("batch_flush_task_%d", atomic.AddUint32(&taskCount, 1)),
		BufferSize:      10,
		IntervalInMills: 5 * time.Second,
		Concurrent:      1,
	}

	option.QueueSize = option.BufferSize * 5
	return option
}

func WithCtx(ctx context.Context) FlushOption {
	return func(option *flushOption) {
		option.ctx = ctx
	}
}

func WithTaskName(name string) FlushOption {
	return func(option *flushOption) {
		option.TaskName = name
	}
}

func WithBufferSize(size int) FlushOption {
	return func(option *flushOption) {
		option.BufferSize = size
	}
}

func WithIntervalInMills(intervalInMills time.Duration) FlushOption {
	return func(option *flushOption) {
		option.IntervalInMills = intervalInMills
	}
}

func WithConcurrent(concurrent int) FlushOption {
	return func(option *flushOption) {
		option.Concurrent = concurrent
	}
}

func WithQueueSize(size int) FlushOption {
	return func(option *flushOption) {
		option.QueueSize = size
	}
}
