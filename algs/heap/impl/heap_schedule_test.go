package impl

import (
	"fmt"
	"testing"
	"time"
)

func TestNewHeapSchedule(t *testing.T) {
	hs := NewHeapSchedule()
	hs.AddJob(&Job{
		timestamp: time.Now().Add(time.Second * 5),
		job: func(data interface{}) {
			fmt.Println("print 5s later")
		},
		data: nil,
	})

	select {}
}
