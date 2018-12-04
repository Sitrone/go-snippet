package example

import (
	"fmt"
	"testing"
	"time"
)

func TestSum(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}
	ret := Sum(s)
	fmt.Println(ret)
}

func TestMuti(t *testing.T) {
	result := fanIn(branch(1), branch(2), branch(3))

	for i := 0; i < 3; i++ {
		fmt.Println(<-result)
	}
}

func TestTimer(t *testing.T) {
	timeout := timer(2 * time.Second) // 定时1s

	for {
		select {
		case <-timeout:
			fmt.Println("already 1s!") // 到时间
			return
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("not timeout, sleep 500 millisecond")
		}
	}
}
