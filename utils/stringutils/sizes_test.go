package stringutils_test

import (
	"testing"

	utils "github.com/daqiancode/gocommons/utils/stringutils"
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
