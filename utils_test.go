package utils

import (
	"testing"

	"gotest.tools/assert"
)

func TestEcho(t *testing.T) {
	assert.Equal(t, Echo("test"), "test")
}

func TestHello(t *testing.T) {
	assert.Equal(t, hello(), "hello")
}
