package example

import (
	"math"
	"testing"
)

func TestInterface(t *testing.T) {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	do(21)
	do("hello")
	do(true)
}

func TestInterfaceCast(t *testing.T) {
	InterfaceCast()

	var tt Tester = FuncDo(func() {
		println("Hello, World!")
	})
	tt.Do()
}
