package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGcd(t *testing.T) {
	gcd(10, 3)
	assert.Equal(t, 1, gcd(10, 3))

	assert.Equal(t, 2, gcd(10, 4))

	assert.Equal(t, 4, gcd(12, 4))

}

func Test_Rotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, nums)
}

func Test_Rotate2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate2(nums, 3)
	assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, nums)
}

func Test_Rotate3(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate3(nums, 3)
	assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, nums)
}

func Test_Reverse(t *testing.T) {
	var nums = []int{1, 2, 3, 4, 5, 6, 7}
	reverse(nums)
	assert.Equal(t, []int{7, 6, 5, 4, 3, 2, 1}, nums)
}
