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

func TestSplitIntoParts(t *testing.T) {
	ids := []string{"999", "1010", "1188511885", "222222", "446446", "38593859", "565656", "824824824", "2121212121"}

	results := map[string][][]string{}
	for _, id := range ids {
		parts := splitIntoParts(id)
		results[id] = parts
	}

	assert.Equal(t, results["999"], [][]string{{"9", "9", "9"}})

	assert.Equal(t, results["1010"], [][]string{
		{"1", "0", "1", "0"},
		{"10", "10"},
	})
	assert.Equal(t, results["1188511885"], [][]string{
		{"1", "1", "8", "8", "5", "1", "1", "8", "8", "5"},
		{"11", "88", "51", "18", "85"},
		{"11885", "11885"},
	})
	assert.Equal(t, results["222222"], [][]string{
		{"2", "2", "2", "2", "2", "2"},
		{"22", "22", "22"},
		{"222", "222"},
	})
	assert.Equal(t, results["446446"], [][]string{
		{"4", "4", "6", "4", "4", "6"},
		{"44", "64", "46"},
		{"446", "446"},
	})
	assert.Equal(t, results["2121212121"], [][]string{
		{"2", "1", "2", "1", "2", "1", "2", "1", "2", "1"},
		{"21", "21", "21", "21", "21"},
		{"21212", "12121"},
	})
}
