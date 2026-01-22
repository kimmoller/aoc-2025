package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneWithTestData(t *testing.T) {
	result, err := Run("test_data", 10)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 40, *result)
}

func TestPartOneWithRealData(t *testing.T) {
	result, err := Run("data", 1000)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 399165, *result)
}

func TestPartTwoWithTestData(t *testing.T) {
	result, err := Run("test_data", 10)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 4277556, *result)
}

func TestPartTwoWithRealData(t *testing.T) {
	result, err := Run("data", 1000)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 8360, *result)
}
