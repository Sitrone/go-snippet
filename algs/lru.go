package algs

import (
	"container/list"
	"fmt"
)

// 其他实现参考github的simplelru
// 慎用interface{}作为入参，丢了类型约束
type LRUCache struct {
	list     *list.List
	m        map[interface{}]*list.Element
	Capacity int
}

func (l *LRUCache) String() string {
	return fmt.Sprintf("list elements:%v, size:%v", l.list, l.Size())
}

func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		Capacity: size,
		list:     list.New(),
		m:        make(map[interface{}]*list.Element, size*5/4+1),
	}
}

type entry struct {
	key, val interface{}
}

func newEntry(key, val interface{}) *entry {
	return &entry{key: key, val: val}
}

func (l *LRUCache) Put(k, val interface{}) {
	if element, ok := l.m[k]; ok {
		l.list.Remove(element)
		delete(l.m, k)
	}

	if len(l.m) >= l.Capacity {
		back := l.list.Back()
		l.list.Remove(back)
		delete(l.m, back.Value.(*entry).key)
	}

	element := l.list.PushFront(newEntry(k, val))
	l.m[k] = element
}

func (l *LRUCache) Get(k interface{}) interface{} {
	if element, ok := l.m[k]; ok {
		l.list.MoveToFront(element)
		return element.Value.(*entry).val
	} else {
		return nil
	}
}

func (l *LRUCache) Size() int {
	return len(l.m)
}

func (l *LRUCache) Remove(k interface{}) {
	if element, ok := l.m[k]; ok {
		l.list.Remove(element)
	}
}