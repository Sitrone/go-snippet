package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	ints := []int{1, 2, 4}
	for k, v := range ints {
		fmt.Println(&k, &v)
	}

	for k, v := range ints[:] {
		fmt.Println(&k, &v)
	}

	for i := 0; i < len(ints); i++ {
		fmt.Println(&ints[i])
	}

	ages := map[string]int{"张三": 15, "李四": 20, "王武": 36}

	for name, age := range ages {
		fmt.Println(&name, &age)
	}
}
