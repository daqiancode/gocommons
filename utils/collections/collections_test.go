package collections_test

import (
	"testing"

	"github.com/daqiancode/gocommons/utils/collections"
	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	a := []int{}
	r := collections.Unique(a)
	assert.Nil(t, r)
	a1 := []int{1, 2, 3, 3, 4, 4}
	r1 := collections.Unique(a1)
	assert.EqualValues(t, []int{1, 2, 3, 4}, r1)
}

func TestIntersection(t *testing.T) {
	a := []int{1, 2, 3, 3, 4, 4}
	b := []int{1}
	r := collections.Intersection(a, b)
	assert.EqualValues(t, []int{1}, r)

}

func TestSubtract(t *testing.T) {
	a := []int{1, 2, 3, 3, 4, 4}
	b := []int{1}
	r := collections.Subtract(a, b)
	assert.EqualValues(t, []int{2, 3, 4}, r)

}

func TestUnion(t *testing.T) {
	a := []int{1, 2, 3, 3, 4, 4}
	b := []int{5, 6}
	r := collections.Union(a, b)
	assert.EqualValues(t, []int{1, 2, 3, 4, 5, 6}, r)
}
