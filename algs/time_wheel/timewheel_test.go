package time_wheel

import (
	"fmt"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tw := New(time.Second, 60)
	tw.AddTask(&Task{
		Key:   "pingTask001",
		Delay: time.Second * 10,
		data:  map[string]string{"hello": "world"},
		Job: func(data interface{}) {
			fmt.Printf("run task ping task 001, data=[%v]\n", data)
		},
	})

	select {}
}
