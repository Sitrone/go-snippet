package flusher

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

const (
	AddCheckBufferSize = 10
)

type Worker struct {
	ctx         context.Context
	taskName    string
	bufferSize  int
	flushTicker *time.Ticker

	addCheckChan chan bool
	bufferChan   chan interface{}
	processor    Processor
}

func newWorker(options flushOption, processor Processor) *Worker {
	return &Worker{
		ctx:         newCtxWithDone(options.ctx),
		taskName:    options.TaskName,
		bufferSize:  options.BufferSize,
		flushTicker: time.NewTicker(options.IntervalInMills),

		addCheckChan: make(chan bool, AddCheckBufferSize),
		bufferChan:   make(chan interface{}, options.QueueSize),
		processor:    processor,
	}
}

func (w *Worker) Add(data interface{}) {
	w.bufferChan <- data
	w.flushOnDemand()
}

func (w *Worker) flushOnDemand() {
	if w.canFlush() {
		w.enableWorker()
	}
}

func (w *Worker) enableWorker() {
	w.addCheckChan <- true
}

func (w *Worker) canFlush() bool {
	return len(w.bufferChan) > w.bufferSize
}

func (w *Worker) run() {
	go func() {
		for {
			select {
			case <-w.flushTicker.C:
				w.flush(false)
			case check := <-w.addCheckChan:
				if check && w.canFlush() {
					w.flush(false)
				}
			case <-w.ctx.Done():
				w.flush(true)
				fmt.Printf("task(%s) interrupted", w.taskName)
				return
			}
		}
	}()
}

func (w *Worker) flush(force bool) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("run flush task (%s) panic!, err=%+v\n", w.taskName, err)
			log.Println(string(debug.Stack()))
		}
	}()

	if len(w.bufferChan) == 0 {
		return
	}

	var (
		curQueueSize = len(w.bufferChan)
		toHandleSize int
	)

	if force {
		toHandleSize = curQueueSize
	} else {
		toHandleSize = min(curQueueSize, w.bufferSize)
	}

	toHandleData := make([]interface{}, 0, toHandleSize)
	for i := 0; i < toHandleSize; i++ {
		toHandleData = append(toHandleData, <-w.bufferChan)
	}

	w.processor(w.ctx, toHandleData)
}

func done() <-chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return c
}

func newCtxWithDone(ctx context.Context) context.Context {
	newCtx, cancel := context.WithCancel(ctx)
	go func() {
		select {
		case s := <-done():
			fmt.Printf("flush task will close, recv:%+v", s)
			cancel()
		}
	}()
	return newCtx
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
