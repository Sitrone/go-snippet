package leetcode

import (
	"fmt"
	"testing"
)

func TestLRUCache_Get(t *testing.T) {
	lruCache := Constructor(2)
	fmt.Println(lruCache.Get(2))
	lruCache.Put(2, 6)
	fmt.Println(lruCache.Get(1))
	lruCache.Put(1, 5)
	lruCache.Put(1, 2)
	fmt.Println(lruCache.Get(1))
	fmt.Println(lruCache.Get(2))

	fmt.Println(10 / 4)
}
