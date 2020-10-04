package flusher

import (
	"errors"
	"sync/atomic"
	"time"
)

type Flusher struct {
	workers []*Worker
	index   uint64
}

func New(processor Processor, flushIntervalInMills time.Duration, batchSize int, options ...FlushOption) *Flusher {
	if flushIntervalInMills < 0 || batchSize < 0 {
		panic(errors.New("illegal parameters"))
	}

	opts := defaultOption()
	opts.IntervalInMills = flushIntervalInMills
	opts.BufferSize = batchSize
	opts.QueueSize = opts.BufferSize * 5 // make sure enough buffer

	for _, o := range options {
		o(opts)
	}

	flushWorkers := make([]*Worker, 0, opts.Concurrent)
	for i := 0; i < opts.Concurrent; i++ {
		flushWorkers = append(flushWorkers, newWorker(*opts, processor))
	}

	for _, worker := range flushWorkers {
		worker.run()
	}

	flusher := &Flusher{
		workers: flushWorkers,
		index:   0,
	}

	return flusher
}

func (f *Flusher) Add(data interface{}) {
	if len(f.workers) == 1 {
		f.workers[0].Add(data)
		return
	}

	atomic.AddUint64(&f.index, 1)
	mod := f.index % uint64(len(f.workers))
	f.workers[mod].Add(data)
}
