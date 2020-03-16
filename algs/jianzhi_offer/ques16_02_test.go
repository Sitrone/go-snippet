package jianzhi_offer

import (
	"fmt"
	"testing"
)

func TestWordsFrequency_Get(t *testing.T) {
	words := []string{"jxf", "c", "c", "qbmgl", "v", "c", "c", "c", "g", "yw", "jxf", "vnsgp", "jxf", "jqpa", "bsgso", "siea", "siea", "c", "jxf", "jxf", "bsgso"}
	f := Constructor(words)
	fmt.Println(f.Get("ywpes"))
}
