package env_test

import (
	"testing"

	"github.com/daqiancode/gocommons/env"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	r := env.Getenv("hello")
	assert.Empty(t, r)
}
