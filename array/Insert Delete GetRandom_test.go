package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	set := Constructor()
	assert.True(t, set.Insert(1))
	//assert.False(t, set.Remove(2))
	//assert.True(t, set.Insert(2))
	//assert.True(t, set.GetRandom() == 1 || set.GetRandom() == 2)
	//assert.True(t, set.Remove(1))
	//assert.Equal(t, 2, set.GetRandom())
}

func TestInsert_false(t *testing.T) {
	set := Constructor()
	set.Insert(1)
	assert.False(t, set.Insert(1))

}

func TestRemove_true(t *testing.T) {
	set := Constructor()
	set.Insert(1)
	assert.True(t, set.Remove(1))
}

func TestRemove_false(t *testing.T) {
	set := Constructor()
	set.Insert(1)
	assert.False(t, set.Remove(2))
}

func TestGetRandom(t *testing.T) {
	set := Constructor()
	set.Insert(1)
	set.Insert(2)
	set.Remove(1)
	assert.Equal(t, 2, set.GetRandom())
}

func TestGetRandom2(t *testing.T) {
	set := Constructor()
	set.Insert(0)
	set.Insert(1)
	assert.Equal(t, set.nums, []int{0, 1})
	set.Remove(0)
	assert.Equal(t, set.nums, []int{1})
	set.Insert(2)
	assert.Equal(t, set.nums, []int{1, 2})
	set.Remove(1)
	assert.Equal(t, 2, set.GetRandom())
}
