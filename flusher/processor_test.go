package flusher

import (
	"context"
	"fmt"
	"sync/atomic"
)

var totalPrint uint32

var PrintProcessor Processor = func(ctx context.Context, datas []interface{}) {
	fmt.Println("start print ", len(datas))
	fmt.Println(datas)

	atomic.AddUint32(&totalPrint, uint32(len(datas)))

	fmt.Println("finished print ", totalPrint)
}
