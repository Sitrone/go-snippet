package leetcode

import "strconv"

// https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/

// 逆波兰式求解
func evalRPN(tokens []string) int {
	var last, lastSecond, ret int
	stack := make([]int, 0, len(tokens)/2)
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			last, lastSecond = stack[len(stack)-1], stack[len(stack)-2]
			if tokens[i] == "+" {
				ret = lastSecond + last
			} else if tokens[i] == "-" {
				ret = lastSecond - last
			} else if tokens[i] == "*" {
				ret = lastSecond * last
			} else if tokens[i] == "/" {
				ret = lastSecond / last
			}

			stack = stack[:len(stack)-2]
			stack = append(stack, ret)
		} else {
			ret, _ = strconv.Atoi(tokens[i])
			stack = append(stack, ret)
		}
	}

	return stack[0]
}
