package example

import (
	"fmt"
	"testing"
)

func TestDaisyChain(t *testing.T)  {
	const max = 100
	nums := Generator() // 初始化一个整数生成器
	number := <-nums    // 从生成器中抓一个整数(2), 作为初始化整数

	for number <= max { // number作为筛子，当筛子超过max的时候结束筛选
		fmt.Println(number)         // 打印素数, 筛子即一个素数
		nums = Filter(nums, number) //筛掉number的倍数
		number = <-nums             // 更新筛子
	}
}
