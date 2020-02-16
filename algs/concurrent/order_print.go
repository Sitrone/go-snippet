package concurrent

import (
	"fmt"
	"time"
)

func printNum(in, out chan struct{}, content int) {
	<-in
	fmt.Printf("%d ", content)
	time.Sleep(100 * time.Millisecond)
	out <- struct{}{}
}
