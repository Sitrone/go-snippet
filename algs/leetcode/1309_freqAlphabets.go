package leetcode

import (
	"bytes"
)

func freqAlphabets(s string) string {
	rs := []byte(s)
	sharp := byte('#')
	m := map[string]byte{
		"1":   'a',
		"2":   'b',
		"3":   'c',
		"4":   'd',
		"5":   'e',
		"6":   'f',
		"7":   'g',
		"8":   'h',
		"9":   'i',
		"10#": 'j',
		"11#": 'k',
		"12#": 'l',
		"13#": 'm',
		"14#": 'n',
		"15#": 'o',
		"16#": 'p',
		"17#": 'q',
		"18#": 'r',
		"19#": 's',
		"20#": 't',
		"21#": 'u',
		"22#": 'v',
		"23#": 'w',
		"24#": 'x',
		"25#": 'y',
		"26#": 'z',
	}

	var (
		buffer bytes.Buffer
	)
	for i := len(rs) - 1; i >= 0; {
		if rs[i] == sharp {
			tmp := []byte{rs[i-2], rs[i-1], rs[i]}
			buffer.WriteByte(m[string(tmp)])
			i -= 3
		} else {
			buffer.WriteByte(m[string(rs[i])])
			i--
		}
	}

	ans := make([]byte, buffer.Len())
	tmpAns := buffer.Bytes()
	for i, j := len(tmpAns)-1, 0; i >= 0 && j < len(tmpAns); i-- {
		ans[j] = tmpAns[i]
		j++
	}
	return string(ans)
}
