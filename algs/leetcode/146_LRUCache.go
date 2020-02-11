package leetcode

import "container/list"

type LRUCache struct {
	m    map[int]*list.Element
	eles *list.List
	cap  int
}

type kv struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:    make(map[int]*list.Element, 10),
		eles: list.New(),
		cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.eles.MoveToFront(v)
		return v.Value.(*kv).val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 存在key值先更新
	if v, ok := this.m[key]; ok {
		this.eles.Remove(v)
		delete(this.m, key)
	}

	if len(this.m) >= this.cap {
		back := this.eles.Back()
		this.eles.Remove(back)
		delete(this.m, back.Value.(*kv).key)
	}

	element := this.eles.PushFront(&kv{
		key: key,
		val: value,
	})
	this.m[key] = element
}
