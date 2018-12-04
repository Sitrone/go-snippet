package example

import (
	"fmt"
	"time"
)

var now = time.Now()

func init() {
	fmt.Printf("now: %v\n", now)
}
func init() {
	fmt.Printf("since: %v\n", time.Now().Sub(now))

	x := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(cap(x[1:3:6]))
}

func testType() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type User struct {
	id   int
	name string
}

func InterfaceCast() {
	var o interface{} = &User{1, "Tom"}

	if i, ok := o.(fmt.Stringer); ok { // ok-idiom
		fmt.Println(i)
	}

	switch v := o.(type) {
	case nil: // o == nil
		fmt.Println("nil")
	case fmt.Stringer: // interface
		fmt.Println(v)
	case func() string: // func
		fmt.Println(v())
	case *User: // *struct
		fmt.Printf("%d, %s\n", v.id, v.name)
	default:
		fmt.Println("unknown")
	}

}

type Tester interface {
	Do()
}
type FuncDo func()

// 某些时候，让函数直接 "实现" 接⼝口能省不少事。
func (f FuncDo) Do() { f() }
