package impl

import "time"

type HeapSchedule struct {
	heap    *JobHeap
	jobChan chan *Job

	closed chan struct{}
}

// 用堆实现的高性能延时任务执行器
func NewHeapSchedule() *HeapSchedule {
	s := &HeapSchedule{
		heap:    &JobHeap{},
		jobChan: make(chan *Job),
		closed:  make(chan struct{}),
	}

	go s.start()

	return s
}

func (s *HeapSchedule) AddJob(j *Job) {
	s.heap.Push(j)
}

func (s *HeapSchedule) start() {
	go func() {
		for {
			if s.heap.Len() == 0 {
				time.Sleep(time.Second * 1)
				continue
			}

			job := s.heap.Pop().(*Job)
			dura := job.timestamp.Sub(time.Now()).Milliseconds()
			if dura >= 0 {
				time.Sleep(time.Millisecond * time.Duration(dura))
			}
			s.jobChan <- job
		}
	}()

	var job *Job
	for {
		select {
		case job = <-s.jobChan:
			go job.job(job.data)
		case <-s.closed:
			return
		}
	}
}

func (s *HeapSchedule) Stop() {
	s.closed <- struct{}{}
}
