package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_5(t *testing.T) {
	profit := maxProfit([]int{7, 1, 5, 3, 6, 4})
	assert.Equal(t, 5, profit)
}

func Test_0(t *testing.T) {
	profit := maxProfit([]int{7, 6, 4, 3, 1})
	assert.Equal(t, 0, profit)
}

func Test_3(t *testing.T) {
	profit := maxProfit([]int{1, 9, 0, 3, 7})
	assert.Equal(t, 8, profit)
}

func Test_8(t *testing.T) {
	profit := maxProfit([]int{2, 9, 0, 3, 8})
	assert.Equal(t, 8, profit)
}
