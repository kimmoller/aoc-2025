package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleValidation(t *testing.T) {
	ids := []string{"999", "1010", "1188511885", "222222", "446446", "38593859", "565656", "824824824", "2121212121"}

	results := simpleValidation(ids)
	assert.Equal(t, []string{"1010", "1188511885", "222222", "446446", "38593859"}, results)
}

func TestComplexValidation(t *testing.T) {
	ids := []string{"999", "1010", "1188511885", "222222", "446446", "38593859", "565656", "824824824", "2121212121", "998", "112111", "121221"}

	result := complexValidation(ids)
	assert.True(t, slices.Contains(result, "999"))
	assert.True(t, slices.Contains(result, "1010"))
	assert.True(t, slices.Contains(result, "1188511885"))
	assert.True(t, slices.Contains(result, "222222"))
	assert.True(t, slices.Contains(result, "446446"))
	assert.True(t, slices.Contains(result, "38593859"))
	assert.True(t, slices.Contains(result, "565656"))
	assert.True(t, slices.Contains(result, "824824824"))
	assert.True(t, slices.Contains(result, "2121212121"))
}
