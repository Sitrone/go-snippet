package algs

import (
	"fmt"
	"testing"
)

func TestDefangIPaddr(t *testing.T) {
	address := "255.100.50.0"
	fmt.Println(DefangIPaddr(address))
}
