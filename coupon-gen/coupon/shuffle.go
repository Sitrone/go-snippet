package coupon

import (
	"math/rand"
	"time"
)

var (
	//rnd = rand.New(&cryptoSource{})
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func shuffle(arr []int) {
	perm := rnd.Perm(len(arr))
	for i, v := range perm {
		arr[i], arr[v] = arr[v], arr[i]
	}
}
