package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test1(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	assert.Equal(t, productExceptSelf(nums), productExceptSelf2(nums))
}
