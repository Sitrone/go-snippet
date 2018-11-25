package main

import "fmt"

/**
defer 栈，后进先出
*/

func main() {
	a()
	b()
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}
