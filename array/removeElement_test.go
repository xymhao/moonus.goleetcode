package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResult_2(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	res := removeElement(nums, val)
	if res != 2 {
		t.Error("res is not 2")
	}
}

func TestResult_5(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	res := removeElement(nums, val)
	if res != 5 {
		t.Error("res is not 5")
	}
	assert.Equal(t, []int{0, 1, 3, 0, 4}, nums[:res])
}

func Test_removeDuplicatesII(t *testing.T) {
	ints := []int{1, 1, 1, 2, 2, 3}
	ii := removeDuplicatesII(ints)
	assert.Equal(t, []int{1, 1, 2, 2, 3}, ints[:ii])

	ints2 := []int{1}
	duplicatesII := removeDuplicatesII(ints2)
	assert.Equal(t, []int{1}, ints2[:duplicatesII])
}

func Test_majorityElement_Boyer_Moore(t *testing.T) {
	moore := majorityElement_Boyer_Moore([]int{2, 3, 3, 2, 3})
	assert.Equal(t, 3, moore)
}
