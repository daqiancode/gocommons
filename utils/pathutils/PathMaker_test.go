package pathutils_test

import (
	"fmt"
	"testing"

	"github.com/daqiancode/gocommons/utils/pathutils"
	"github.com/stretchr/testify/assert"
)

func TestPathMaker(t *testing.T) {
	p1 := pathutils.MakeDateRandPath("jpg")
	p2 := pathutils.MakeIdPath("12", "jpg")
	p3 := pathutils.MakeRandPath("jpg")
	assert.Equal(t, "00/12/0012.jpg", p2)
	assert.Equal(t, 30, len(p3))
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

}
