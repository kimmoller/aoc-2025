package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneWithTestData(t *testing.T) {
	result, err := Day5("day_5_test_data")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 3, *result)
}

func TestPartOneWithRealData(t *testing.T) {
	result, err := Day5("day_5_data")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 611, *result)
}

func TestPartTwoWithTestData(t *testing.T) {
	result, err := Day5("day_5_test_data")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 3, *result)
}

func TestPartTwoWithRealData(t *testing.T) {
	result, err := Day5("day_5_data")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 3, *result)
}
