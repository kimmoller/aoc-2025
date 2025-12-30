package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneWithTestData(t *testing.T) {
	result, err := Run("test_data", false)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 3, *result)
}

func TestPartOneWithRealData(t *testing.T) {
	result, err := Run("data", false)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 611, *result)
}

func TestPartTwoWithTestData(t *testing.T) {
	result, err := Run("test_data", true)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 14, *result)
}

func TestPartTwoWithRealData(t *testing.T) {
	result, err := Run("data", true)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 3, *result)
}
