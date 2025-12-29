package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneWithTestData(t *testing.T) {
	result, err := Day4("day_4_test_data", 2)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 13, *result)
}

func TestPartOneWithRealData(t *testing.T) {
	result, err := Day4("day_4_data", 2)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 1349, *result)
}

// FIX: Copied from Day 3
func TestPartTwoWithTestData(t *testing.T) {
	result, err := Day4("day_4_test_data", 12)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 3121910778619, *result)
}

// FIX: Copied from Day 3
func TestPartTwoWithRealData(t *testing.T) {
	result, err := Day4("day_4_data", 12)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 169935154100102, *result)
}
