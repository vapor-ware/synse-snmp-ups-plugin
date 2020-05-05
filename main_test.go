package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Basic test to ensure that plugin creation is valid and all plugin
// specific components are registered successfully.
func TestMakePlugin(t *testing.T) {
	_, err := MakePlugin()
	assert.NoError(t, err)
}
