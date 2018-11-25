package main

import "fmt"

type Vertex struct {
	x int
	y int
}

/**
 * 与 C 不同，Go 没有指针运算。
 */
func main() {
	i, j := 42, 2701

	p := &i // point to i
	fmt.Println(p)
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	v := Vertex{1, 2}
	fmt.Println(v.x)
	fmt.Println(v.y)
}
