package hard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_5(t *testing.T) {
	assert.Equal(t, 5, candy([]int{1, 0, 2}))
}
