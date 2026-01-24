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
	assert.Equal(t, float64(50), *result)
}

func TestPartOneWithRealData(t *testing.T) {
	result, err := Run("data", false)
	if err != nil {
		panic(err)
	}
	intResult := int(*result)
	assert.Equal(t, 4759531084, intResult)
}

func TestPartTwoWithTestData(t *testing.T) {
	result, err := Run("test_data", true)
	if err != nil {
		panic(err)
	}
	intResult := int(*result)
	assert.Equal(t, 24, intResult)
}

func TestPartTwoWithRealData(t *testing.T) {
	result, err := Run("data", true)
	if err != nil {
		panic(err)
	}
	intResult := int(*result)
	assert.Equal(t, 1539238860, intResult)
}
