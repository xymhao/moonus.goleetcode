package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hIndex(t *testing.T) {
	assert.Equal(t, 3, hIndex([]int{0, 1, 3, 5, 6}))
	assert.Equal(t, 1, hIndex([]int{100}))
}

func Test_hIndex2(t *testing.T) {
	assert.Equal(t, 3, hIndex2([]int{0, 1, 3, 5, 6}))
	assert.Equal(t, 1, hIndex2([]int{100}))
	assert.Equal(t, 1, hIndex2([]int{1, 3, 1}))
}
