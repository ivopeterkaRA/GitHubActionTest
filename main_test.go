package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func Test_getMessage(t *testing.T) {
	assert_ := assert.New(t)
	result := getTheMessage()
	assert_.Equal("Hello World", result)
}
