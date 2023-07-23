package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResult_True(t *testing.T) {
	jump := canJump([]int{2, 3, 1, 1, 4})
	assert.True(t, jump)

	assert.True(t, canJump([]int{3, 0, 8, 2, 0, 0, 1}))
}

func Test_False(t *testing.T) {
	jump := canJump([]int{3, 2, 1, 0, 4})
	assert.False(t, jump)

	assert.False(t, canJump([]int{3, 2, 1, 0, 4}))
}

func Test_jump_2(t *testing.T) {
	assert.Equal(t, 2, jump([]int{2, 3, 1, 1, 4}))
	assert.Equal(t, 2, jump([]int{2, 3, 0, 1, 4}))
}

func Test_jump2(t *testing.T) {
	assert.Equal(t, 2, jump2([]int{2, 3, 1, 1, 4}))
	assert.Equal(t, 2, jump2([]int{2, 3, 0, 1, 4}))
}
