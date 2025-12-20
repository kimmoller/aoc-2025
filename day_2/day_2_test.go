package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartOneWithTestData(t *testing.T) {
	sum, err := Day2("day_2_test_data", SIMPLE)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1227775554, *sum)
}

func TestPartOneWithRealData(t *testing.T) {
	sum, err := Day2("day_2_data", SIMPLE)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 32976912643, *sum)
}

func TestPartTwoWithTestData(t *testing.T) {
	sum, err := Day2("day_2_test_data", COMPLEX)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 4174379265, *sum)
}

func TestPartTwoWithRealData(t *testing.T) {
	sum, err := Day2("day_2_data", COMPLEX)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 54446379122, *sum)
}
