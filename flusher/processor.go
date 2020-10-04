package flusher

import "context"

type Processor func(ctx context.Context, datas []interface{})
