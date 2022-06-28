package randutils_test

import (
	"testing"

	"github.com/daqiancode/gocommons/utils/randutils"
	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {
	assert.NotEqual(t, randutils.RandomAlphabetNumber(32), randutils.RandomAlphabetNumber(32))
	assert.NotEqual(t, randutils.RandomHex(32), randutils.RandomHex(32))
	assert.NotEqual(t, randutils.RandomAlphabetNumberLower(32), randutils.RandomAlphabetNumberLower(32))
	assert.NotEqual(t, randutils.RandomNumber(32), randutils.RandomNumber(32))
}
