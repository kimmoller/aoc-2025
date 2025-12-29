package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneWithTestData(t *testing.T) {
	result, err := Day3("day_3_test_data", 2)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 357, *result)
}

func TestPartOneWithRealData(t *testing.T) {
	result, err := Day3("day_3_data", 2)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 17142, *result)
}

func TestPartTwoWithTestData(t *testing.T) {
	result, err := Day3("day_3_test_data", 12)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 3121910778619, *result)
}

func TestPartTwoWithRealData(t *testing.T) {
	result, err := Day3("day_3_data", 12)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 169935154100102, *result)
}
