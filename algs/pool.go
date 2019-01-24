package algs

import (
	"fmt"
	"sync"
)

type worker struct {
	Func func()
}

// 更加全面的pool参考fasthttp的pool
func GoPoll() {
	var wg sync.WaitGroup

	workers := make(chan worker)
	// 启动一定数量的goroutine，通过闭包传递执行函数
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for wk := range workers {
				wk.Func()
			}
		}()
	}

	for i := 0; i < 100; i++ {
		j := i
		wk := worker{
			Func: func() {
				fmt.Println(i + j)
			},
		}

		workers <- wk
	}

	close(workers)
	wg.Wait()
}