package jianzhi_offer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberOf1(t *testing.T) {
	var n = 10
	assert.Equal(t, NumberOf1(n), 2)
}
