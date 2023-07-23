package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canCompleteCircuit(t *testing.T) {
	assert.Equal(t, 3, canCompleteCircuit2([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	assert.Equal(t, -1, canCompleteCircuit2([]int{2, 3, 4}, []int{3, 4, 3}))
}

func Test_canCompleteCircuit_4(t *testing.T) {
	assert.Equal(t, 4, canCompleteCircuit2([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}))
}
func Test_canCompleteCircuit_3(t *testing.T) {
	assert.Equal(t, 3, canCompleteCircuit2([]int{5, 8, 2, 8}, []int{6, 5, 6, 6}))
}

func Test_canCompleteCircuit_(t *testing.T) {
	assert.Equal(t, -1, canCompleteCircuit2([]int{2, 3, 4}, []int{3, 4, 3}))
}
