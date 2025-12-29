package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneWithTestData(t *testing.T) {
	result, err := Day6("day_6_test_data")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 4277556, *result)
}

func TestPartOneWithRealData(t *testing.T) {
	result, err := Day6("day_6_data")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 4583860641327, *result)
}

func TestPartTwoWithTestData(t *testing.T) {
	result, err := Day6("day_6_test_data")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 4277556, *result)
}

func TestPartTwoWithRealData(t *testing.T) {
	result, err := Day6("day_6_data")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 4277556, *result)
}
