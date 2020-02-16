package concurrent

import (
	"testing"
	"time"
)

// 交替打印1，2，3
func TestOrderPrint(t *testing.T) {
	chanA, chanB, chanC := make(chan struct{}), make(chan struct{}), make(chan struct{})

	go func() {
		for {
			printNum(chanA, chanB, 1)
		}
	}()
	go func() {
		for {
			printNum(chanB, chanC, 2)
		}
	}()
	go func() {
		for {
			printNum(chanC, chanA, 3)
		}
	}()

	// 启动
	chanA <- struct{}{}

	time.Sleep(time.Second * 1)
}
