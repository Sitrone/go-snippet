package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestIntHeap_Len(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println()
}

func TestPriorityQueue_Len(t *testing.T) {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
}
