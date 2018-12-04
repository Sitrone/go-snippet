package example

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (i IPAddr)String() string {
	s := make([]string, len(i))
	for e := range i {
		fmt.Println(int(i[e]))
		s[e] = strconv.Itoa(int(i[e]))
	}
	return fmt.Sprintf("ipAddress is:%v", strings.Join(s, "."))
	//return fmt.Sprintf("%d.%d.%d.%d", p[0], p[1], p[2], p[3])
}
//func main() {
//	hosts := map[string]IPAddr{
//		"loopback":  {127, 0, 0, 1},
//		"googleDNS": {8, 8, 8, 8},
//	}
//	for name, ip := range hosts {
//		fmt.Printf("%v: %v\n", name, ip)
//	}
//}
