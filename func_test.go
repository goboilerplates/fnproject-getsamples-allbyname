package main

import (
	"bufio"
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHandle .
func TestHandle(test *testing.T) {
	reader := strings.NewReader("test")
	writer := bufio.NewWriter(&bytes.Buffer{})
	HandleRequest(context.Background(), reader, writer)

	assert.True(test, TestResult)
}

// TestCreateAPI .
func TestCreateAPI(test *testing.T) {
	result := CreateAPI()
	assert.NotNil(test, result)
}
