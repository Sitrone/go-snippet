package algs

import (
	"fmt"
	"testing"
)

func TestNewLRUCache(t *testing.T) {
	lruCache := NewLRUCache(3)
	lruCache.Put(1, "1")
	lruCache.Put(2, "2")
	lruCache.Put(3, "3")
	fmt.Println(lruCache.Get(1))
	fmt.Println(lruCache)
	lruCache.Put(4, "4")
	lruCache.Put(5, "5")
	fmt.Println(lruCache.Get(5))
	fmt.Println(lruCache)

	s := "m.y+name@email.com"
	for _, v := range s {
		fmt.Println(v == 'm')
	}
}
