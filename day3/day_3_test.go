package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneWithTestData(t *testing.T) {
	result, err := Day3("day_3_test_data")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 357, *result)
}

func TestPartOneWithRealData(t *testing.T) {
	result, err := Day3("day_3_data")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 17142, *result)
}
