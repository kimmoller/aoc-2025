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
	assert.Equal(t, 171503, *result)
}

func TestPartTwoWithTestData(t *testing.T) {
	result, err := RunLimitless("test_data")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 25272, *result)
}

func TestPartTwoWithRealData(t *testing.T) {
	result, err := RunLimitless("data")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 9069509600, *result)
}
