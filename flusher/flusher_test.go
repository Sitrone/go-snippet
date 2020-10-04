package flusher

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFlusher_Add(t *testing.T) {
	flusher := New(PrintProcessor, 2*time.Second, 10,
		WithTaskName("flush_user_login_log_tab"),
		WithConcurrent(2),
	)

	for i := 0; i < 200; i++ {
		flusher.Add(fmt.Sprintf("hello task %d", i))
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	}

	time.Sleep(time.Second * 5)
}
