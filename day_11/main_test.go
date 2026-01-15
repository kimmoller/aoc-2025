package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneWithTestData(t *testing.T) {
	result, err := Run("test_data", "you", "out", false)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 5, *result)
}

func TestPartOneWithRealData(t *testing.T) {
	result, err := Run("data", "you", "out", false)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 590, *result)
}

func TestPartTwoWithTestData(t *testing.T) {
	result, err := Run("test_data_2", "svr", "out", true)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 2, *result)
}

func TestPartTwoWithRealData(t *testing.T) {
	result, err := Run("data", "svr", "out", true)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 4277556, *result)
}

func TestSumBasedPartTwoWithTestData(t *testing.T) {
	result, err := RunSumBased("test_data_2", "svr", "out", true)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, int64(2), *result)
}

func TestSumBasedPartTwoWithRealData(t *testing.T) {
	result, err := RunSumBased("data", "svr", "out", true)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, int64(553204221431080), *result)
}
