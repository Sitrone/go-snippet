package cryptor

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRandString(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandString(6))
	}
}

func TestCryptoRandSource_Int63(t *testing.T) {
	r := rand.New(NewCryptoRandSource())

	letters := []rune(defaultCharacters)

	for i := 0; i < 10; i++ {
		dest := make([]rune, 6)
		perm := r.Perm(len(defaultCharacters))
		for i := range dest {
			dest[i] = letters[perm[i]]
		}
		fmt.Println(string(dest))
	}
}
