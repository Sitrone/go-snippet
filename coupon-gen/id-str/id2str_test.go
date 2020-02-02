package id_str

import (
	"fmt"
	"testing"
)

func TestNewIdStr(t *testing.T) {
	ids := NewIdStr()
	str := ids.Id2String(1_233_154_314)
	fmt.Println(str)

	id := ids.Str2Id(str)
	fmt.Println(id)

	fmt.Println(ids.MaxId())
}
