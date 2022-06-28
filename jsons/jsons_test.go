package jsons_test

import (
	"testing"
	"time"

	"github.com/daqiancode/gocommons/jsons"
	"github.com/stretchr/testify/assert"
)

type A struct {
	Date time.Time
}

func TestJSON(t *testing.T) {
	a := A{Date: time.Now()}
	s, err := jsons.JSON.MarshalToString(a)
	assert.Nil(t, err)
	var b A
	jsons.JSON.UnmarshalFromString(s, &b)
	assert.True(t, b.Date.Equal(a.Date))
}
