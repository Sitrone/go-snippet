package example

import (
	"math/rand"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func Sum(s []int) int {
	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	// 使用此种方式，无法控制go程的调度，因此无法控制程序的退出
	//quit := make(chan int)
	//go func() { quit <- 1 }()
	// https://stackoverflow.com/questions/8593645/is-it-ok-to-leave-a-channel-open
	//defer close(c)
	//defer close(quit)
	//var sum int
	//for isQuit := false; !isQuit; {
	//	select {
	//	case v := <-c:
	//		sum += v
	//	case <-quit:
	//		isQuit = true
	//	}
	//}
	//return sum
	x, y := <-c, <-c // 从 c 中接收
	return x + y
}

// 一个比较耗时的事情，比如计算
func doStuff(x int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second) //模拟计算
	return 100 - x                                         // 假如100-x是一个很费时的计算
}

// 每个分支开出一个goroutine做计算并把计算结果流入各自信道
func branch(x int) chan int {
	ch := make(chan int)
	go func() { ch <- doStuff(x) }()
	return ch
}

//func fanIn(chs ... chan int) chan int {
//	ch := make(chan int)
//
//	for _, c := range chs {
//		// 注意此处明确传值
//		go func(c chan int) {
//			ch <- <-c
//		}(c) // 复合
//	}
//
//	return ch
//}

func fanIn(branches ... chan int) chan int {
	c := make(chan int)

	go func() {
		//select会尝试着依次取出各个信道的数据
		for i := 0; i < len(branches); i++ {
			select {
			case v := <-branches[i]:
				c <- v
			}
		}
	}()

	return c
}

func timer(duration time.Duration) <-chan time.Time {
	var timeOut <-chan time.Time
	timeOut = time.After(duration)

	//go func() {
	//	time.Sleep(duration)
	//	ch <- true // 到时间啦！
	//}()

	return timeOut
}
