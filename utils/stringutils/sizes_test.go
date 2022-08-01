package utils_test

import (
	"testing"

	"github.com/bitcrycode/product/core/common/utils"
	"github.com/stretchr/testify/assert"
)

func TestSizes(t *testing.T) {
	r, err := utils.ParseSize("80x80")
	assert.Nil(t, err)
	assert.Equal(t, []int{80, 80}, r)
	r1, err := utils.ParseSize("80")
	assert.Nil(t, err)
	assert.Equal(t, []int{80, 80}, r1)
}
