package main

import (
	"testing"

	"github.com/goboilerplates/core"
	"github.com/stretchr/testify/assert"
)

// TestGetSamplesAllByNameAPIImplAll .
func TestGetSamplesAllByNameAPIImplAll(test *testing.T) {
	api := GetSamplesAllByNameAPIImpl{
		Interactor: core.CreateDefaultGetSamples("Hello"),
	}

	assert.NotNil(test, api)
	assert.NotNil(test, api.Interactor)

	result, err := api.AllByName("h")
	assert.Nil(test, err)

	assert.NotNil(test, result)
	assert.True(test, len(result) > 0)
}
