package gotour

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

//在内部，接口值可以看做包含值和具体类型的元组：
//
//(value, type)
//接口值保存了一个具体底层类型的具体值。
//
//接口值调用方法时会执行其底层类型的同名方法。
func testInterface() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}
