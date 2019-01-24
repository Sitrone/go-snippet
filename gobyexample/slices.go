package main

import "fmt"

// 理解slice的append的扩容机制
// 元素小于1024个的时候，双倍扩容，大于1024的时候增加1/4

// https://www.zhihu.com/question/27161493
// https://blog.golang.org/go-slices-usage-and-internals

func testSlice1()  {
	s := []int{5}
	s = append(s, 7)
	s = append(s, 9)
	x := append(s, 11)
	y := append(s, 12)
	fmt.Println(s, x, y)
	// [5 7 9] [5 7 9 12] [5 7 9 12]
}

func testSlice2() {
	s := []int{5, 7, 9}
	x := append(s, 11)
	y := append(s, 12)
	fmt.Println(s,x,y)

	//[5 7 9] [5 7 9 11] [5 7 9 12]
}
