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
	assert.Equal(t, 4277556, *result)
}

func TestPartOneWithRealData(t *testing.T) {
	result, err := Run("data", false)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 4583860641327, *result)
}

func TestPartTwoWithTestData(t *testing.T) {
	result, err := Run("test_data", true)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 3263827, *result)
}

func TestPartTwoWithRealData(t *testing.T) {
	result, err := Run("data", true)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 11602774058280, *result)
}
