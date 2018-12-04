package example

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	const times = 100
	const length = 100
	expect := make([]int, length)
	original := make([]int, length)
	ass := require.New(t)
	for i := 0; i < times; i++ {
		nums := rand.Perm(length)
		copy(expect, nums)
		copy(original, nums)

		sort.Ints(expect)
		Qsort(original)

		ass.Equal(expect, original, "%v", nums)
	}
}
