package example

import (
	"fmt"
)

type Tree struct {
	value int
	left  *Tree
	right *Tree
}

//func main() {
//	const pi float64 = 3.14
//	fmt.Print("Hello world.")
//
//	t := test(5)
//	t()
//
//	testSlice()
//
//	tree := Tree{1, nil, nil}
//	fmt.Println(tree.left, tree.right)
//
//	f := fibonacci()
//	for i := 0; i < 10; i++ {
//		fmt.Println(f())
//	}
//}

func test(x int) func() {
	return func() {
		fmt.Println(x)
	}
}

func testDefer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func testSlice() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	var arr = [10]int{1, 2, 3, 4, 5, 6}
	sl := arr[2:5] //创建有3个元素的slice
	fmt.Println(sl)
	fmt.Println(len(sl))
	fmt.Println(cap(sl))

	b := make([]int, 0, 1)
	ints := append(b, 1, 2, 3)
	fmt.Println(ints)
}

//Go 没有类。不过你可以为结构体类型定义方法。
//
//方法就是一类带特殊的 接收者 参数的函数。
func testPointer() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		temp := a
		a, b = b, a + b
		return temp
	}
}
