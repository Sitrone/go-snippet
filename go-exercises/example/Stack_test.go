package example

import (
	"fmt"
	"testing"
	"time"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Pop())
	stack.Push(3)
	fmt.Println(stack)
}

func TestStack_IsEmpty(t *testing.T) {
	task := &RuleTask{
		updateCycle: make(chan time.Duration),
		ticker:      time.NewTicker(time.Second * 5),
		closed:      make(chan struct{}),
		lastTime:    time.Now().Unix(),
	}

	task.f = func() {
		now := time.Now().Unix()
		fmt.Printf("task duration=%d\n", now-task.lastTime)
		task.lastTime = now
	}

	go task.Run()

	time.Sleep(time.Second * 10)
	task.updateCycle <- time.Second * 2

	time.Sleep(time.Second * 10)
	close(task.closed)
	fmt.Println("over")
}

type RuleTask struct {
	lastTime    int64
	updateCycle chan time.Duration
	ticker      *time.Ticker

	closed chan struct{}

	f func()
}

func (r *RuleTask) Run() {
	r.f()

	for {
		select {
		case <-r.ticker.C:
			r.f()
		case cycle := <-r.updateCycle:
			r.ticker = time.NewTicker(cycle)
		case <-r.closed:
			fmt.Println("task closed")
			return
		}
	}
}
