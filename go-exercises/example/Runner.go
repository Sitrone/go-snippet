package example

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

/**
Ref :http://www.flysnow.org/2017/04/29/go-in-action-go-runner.html
 */
var ErrTimeOut = errors.New("执行者执行超时")
var ErrInterrupt = errors.New("执行者被中断")

type Runner struct {
	tasks     []func(int)
	complete  chan error
	timeout   <-chan time.Time
	interrupt chan os.Signal
}

func NewRunner(tm time.Duration) *Runner {
	return &Runner{
		// 它必须是同步通道，要让main routine等待，一直要任务完成或者被强制终止
		complete:  make(chan error),
		timeout:   time.After(tm),
		// 加缓冲通道，Go runtime在发送这个信号的时候不会被阻塞
		interrupt: make(chan os.Signal, 1),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// 执行任务，执行的过程中接收到中断信号时，返回中断错误
// 如果任务全部执行完，还没有接收到中断信号，则返回nil
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.isInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

//检查是否接收到了中断信号
func (r *Runner) isInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

//开始执行所有任务，并且监视通道事件
func (r *Runner) Start() error {
	//希望接收哪些系统信号
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}
