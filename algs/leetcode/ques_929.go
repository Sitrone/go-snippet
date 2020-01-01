package leetcode

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	m := make(map[string]struct{})
	for _, v := range emails {
		m[formatEmail1(v)] = struct{}{}
	}
	fmt.Println(m)
	return len(m)
}

// 单次遍历，速度最快，但是判断比较多容易出错
// test.email+alex@leetcode.com
func formatEmail(email string) string {
	var temp = strings.Builder{}
	temp.Grow(len(email))
	meetAt := false
	meetPlus := false
	for _, v := range email {
		if v == '@' {
			meetAt = true
			temp.WriteRune(rune(v))
		} else if v == '+' {
			meetPlus = true
		} else if v == '.' && !meetAt {
			continue
		} else if meetPlus && !meetAt {
			continue
		} else {
			temp.WriteRune(rune(v))
		}
	}

	return temp.String()
}

// 多次遍历
func formatEmail1(email string) string {
	i := strings.Index(email, "@")
	local := email[:i]
	rest := email[i:]
	if strings.Contains(local, "+") {
		local = local[:strings.Index(local, "+")]
	}
	local = strings.Replace(local, ".", "", -1)
	return local + rest
}
