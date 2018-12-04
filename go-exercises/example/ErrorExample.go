package example

import (
	"fmt"
	"math"
)

// 定义类型
type ErrNegativeSqrt float64

// 重写Error()
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:  %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

//func main() {
//	//通常函数会返回一个 error 值，调用的它的代码应当判断这个错误是否等于 nil，    来进行错误处理。
//	//这里只是简单的打印
//	fmt.Println(Sqrt(2))
//	fmt.Println(Sqrt(-2))
//}
