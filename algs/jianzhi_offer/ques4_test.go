package jianzhi_offer

import (
	"fmt"
	"testing"
)

func TestReplaceBlank(t *testing.T) {
	s := "Are you OK?"
	fmt.Println(ReplaceBlank(s))

	fmt.Println(ReplaceBlank1(s))
}
