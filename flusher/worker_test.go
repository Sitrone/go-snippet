package flusher

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFlushWorker(t *testing.T) {
	option := defaultOption()
	option.IntervalInMills = 2 * time.Second
	flushWorker := newWorker(*option, PrintProcessor)
	flushWorker.run()

	for i := 0; i < 100; i++ {
		flushWorker.Add(fmt.Sprintf("hello task %d", i))
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}

	time.Sleep(option.IntervalInMills)
}
