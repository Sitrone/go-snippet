package algs

// ref: https://github.com/eachain/aca
// 多模匹配算法 TODO
type aca struct {
}

func DefangIPaddr(address string) string {
	var d = []rune("[.]")
	var ret []rune

	for _, v := range []rune(address) {
		if v == '.' {
			ret = append(ret, d...)
		} else {
			ret = append(ret, v)
		}
	}
	return string(ret)
}
