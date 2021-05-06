package time_wheel

import (
	"container/list"
	"time"
)

type Job func(interface{})

type Task struct {
	Key   string // 唯一标识
	Delay time.Duration
	data  interface{}
	Job
}

type Position struct {
	slot, circle int
}

// TimeWheel TODO 多级时间轮
type TimeWheel struct {
	interval time.Duration
	ticker   *time.Ticker
	slots    []map[int]*list.List
	slotsNum int
	curStep  int

	timerM map[string]*Position

	addTaskChan chan *Task
	delTaskChan chan string
	closed      chan struct{}
}

// New new一个时间轮
func New(interval time.Duration, slotsNum int) *TimeWheel {
	slots := make([]map[int]*list.List, slotsNum)
	for i := 0; i < len(slots); i++ {
		slots[i] = make(map[int]*list.List, 0)
	}

	tw := &TimeWheel{
		interval:    interval,
		ticker:      time.NewTicker(interval),
		slots:       slots,
		slotsNum:    slotsNum,
		curStep:     -1,
		timerM:      make(map[string]*Position, slotsNum),
		addTaskChan: make(chan *Task),
		delTaskChan: make(chan string),
		closed:      make(chan struct{}),
	}

	go tw.start()

	return tw
}

func (t *TimeWheel) AddTask(task *Task) {
	t.addTaskChan <- task
}

func (t *TimeWheel) DelTask(key string) {
	if key == "" {
		return
	}

	t.delTaskChan <- key
}

func (t *TimeWheel) Stop() {
	close(t.closed)
}

func (t *TimeWheel) start() {
	for {
		select {
		case <-t.ticker.C:
			//fmt.Println("scan one slot")
			t.handleTicker()
		case task := <-t.addTaskChan:
			t.addTask(task)
		case key := <-t.delTaskChan:
			t.delTask(key)
		case <-t.closed:
			t.ticker.Stop()
			return
		}
	}

}

func (t *TimeWheel) handleTicker() {
	t.curStep++
	slot, curCycle := t.curStep%t.slotsNum, t.curStep/t.slotsNum

	// 处理当前circle
	if l, ok := t.slots[slot][curCycle]; ok {
		if l != nil && l.Len() != 0 {
			for e := l.Front(); e != nil; {
				task := e.Value.(*Task)
				go task.Job(task.data)

				next := e.Next()
				l.Remove(e)
				if task.Key != "" {
					delete(t.timerM, task.Key)
				}

				e = next
			}
		}
	}
}

func (t *TimeWheel) delTask(key string) {
	pos, ok := t.timerM[key]
	if !ok {
		return
	}

	m := t.slots[pos.slot]
	if len(m) == 0 {
		return
	}
	l := m[pos.circle]
	if l == nil || l.Len() == 0 {
		return
	}

	for e := l.Front(); e != nil; {
		task := e.Value.(*Task)
		//log.Printf("remove delay task, key=%s", key)
		if task.Key == key {
			delete(t.timerM, task.Key)
			l.Remove(e)
		}

		e = e.Next()
	}
}

func (t *TimeWheel) addTask(task *Task) {
	pos := t.getPosition(task.Delay)
	if t.slots[pos.slot][pos.circle] == nil {
		t.slots[pos.slot][pos.circle] = list.New()
	}

	t.slots[pos.slot][pos.circle].PushBack(task)
	if task.Key != "" {
		t.timerM[task.Key] = pos
	}
}

func (t *TimeWheel) getPosition(delay time.Duration) *Position {
	if delay < t.interval {
		delay = t.interval
	}

	steps := int(delay / t.interval)
	slot, circle := (t.curStep+steps)%t.slotsNum, (steps-1)/t.slotsNum
	return &Position{
		slot:   slot,
		circle: circle,
	}
}
