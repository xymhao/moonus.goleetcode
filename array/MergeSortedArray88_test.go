package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_88_merge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3

	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, nums1)
}

func Test_1(t *testing.T) {
	nums1 := []int{1}
	m := 1

	nums2 := []int{}
	n := 0
	merge(nums1, m, nums2, n)
	assert.Equal(t, []int{1}, nums1)
}

func Test_2(t *testing.T) {
	nums1 := []int{0}
	m := 0

	nums2 := []int{1}
	n := 1
	merge(nums1, m, nums2, n)
	assert.Equal(t, []int{1}, nums1)
}

func Test_Move2(t *testing.T) {
	nums1 := []int{1, 3, 5, 7, 9, 0, 0, 0, 0, 0}
	m := 5

	nums2 := []int{2, 4, 6, 8, 10}
	n := 5
	merge(nums1, m, nums2, n)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nums1)
}

func Test_Move3(t *testing.T) {
	nums1 := []int{2, 4, 6, 8, 10, 0, 0, 0, 0, 0}
	n := 5

	nums2 := []int{1, 3, 5, 7, 9}
	m := 5
	merge(nums1, m, nums2, n)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, nums1)
}

func Test4(t *testing.T) {
	nums1 := []int{1, 2, 4, 5, 6, 0}
	m := 5

	nums2 := []int{3}
	n := 1

	merge(nums1, m, nums2, n)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, nums1)
}

func Test5(t *testing.T) {
	nums1 := []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	m := 6

	nums2 := []int{1, 2, 3}
	n := 3

	merge(nums1, m, nums2, n)
	assert.Equal(t, []int{-1, 0, 0, 1, 2, 3, 3, 3, 3}, nums1)
}
