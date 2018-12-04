package example

import (
	"fmt"
	"sort"
	"testing"
)

func TestGetSwap(t *testing.T) {
	ret := GetSwap()
	s := make([]SwapView, len(ret))
	for e := range ret {
		s = append(s, *e)
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i].Size < s[j].Size
	})

	fmt.Printf("%5s %9s %s\n", "PID", "SWAP", "COMMAND")
	var total int64
	for _, v := range s {
		fmt.Printf("%5d %9s %s\n", v.Pid, FormatSize(v.Size), v.Comm)
		total += v.Size
	}
	fmt.Printf("Total: %8s\n", FormatSize(total))
}

func TestFormatSize(t *testing.T) {
	fmt.Println(FormatSize(1024))
}
