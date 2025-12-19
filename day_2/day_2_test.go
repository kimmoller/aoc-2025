package main

import (
	"testing"

	"aoc2025/utils"

	"github.com/stretchr/testify/assert"
)

func TestPartOneWithTestData(t *testing.T) {
	data, err := utils.ReadData("day_2_test_data")
	if err != nil {
		panic(err)
	}

	sum, err := SumOfInvalidIds(data, false)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1227775554, *sum)
}

func TestPartOneWithRealData(t *testing.T) {
	data, err := utils.ReadData("day_2_data")
	if err != nil {
		panic(err)
	}

	sum, err := SumOfInvalidIds(data, false)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 32976912643, *sum)
}

func TestPartTwoWithTestData(t *testing.T) {
	data, err := utils.ReadData("day_2_test_data")
	if err != nil {
		panic(err)
	}

	sum, err := SumOfInvalidIds(data, true)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 4174379265, *sum)
}
