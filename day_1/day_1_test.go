package main

import (
	"testing"

	"aoc2025/utils"

	"github.com/stretchr/testify/assert"
)

func TestWithTestData(t *testing.T) {
	data, err := utils.ReadData("day_1_test_data")
	if err != nil {
		panic(err)
	}

	password, err := PassowrdFromDial(data, false)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 3, *password)
}

func TestWithRealData(t *testing.T) {
	data, err := utils.ReadData("day_1_data")
	if err != nil {
		panic(err)
	}

	password, err := PassowrdFromDial(data, false)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1034, *password)
}

func TestPartTwoWithTestData(t *testing.T) {
	data, err := utils.ReadData("day_1_test_data")
	if err != nil {
		panic(err)
	}

	password, err := PassowrdFromDial(data, true)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 6, *password)
}

func TestPartTwoWithRealData(t *testing.T) {
	data, err := utils.ReadData("day_1_data")
	if err != nil {
		panic(err)
	}

	password, err := PassowrdFromDial(data, true)
	if err != nil {
		panic(err)
	}

	assert.Greater(t, *password, 5368)
	assert.Less(t, *password, 6173)
	assert.Equal(t, 6166, *password)
}
