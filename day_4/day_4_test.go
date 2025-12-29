package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneWithTestData(t *testing.T) {
	result, err := Day4("day_4_test_data", false)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 13, *result)
}

func TestPartOneWithRealData(t *testing.T) {
	result, err := Day4("day_4_data", false)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 1349, *result)
}

func TestPartTwoWithTestData(t *testing.T) {
	result, err := Day4("day_4_test_data", true)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 43, *result)
}

func TestPartTwoWithRealData(t *testing.T) {
	result, err := Day4("day_4_data", true)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 8277, *result)
}
