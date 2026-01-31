package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestAllCombinations(t *testing.T) {
	allCombinations := map[int][][][]int{
		2: {
			{{0, 1}},
		},

		3: {
			{{0, 1}, {0, 2}, {1, 2}},
			{{0, 1, 2}},
		},

		4: {
			{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}},
			{{0, 1, 2}, {0, 1, 3}, {0, 2, 3}, {1, 2, 3}},
			{{0, 1, 2, 3}},
		},

		5: {
			{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
			{{0, 1, 2}, {0, 1, 3}, {0, 1, 4}, {0, 2, 3}, {0, 2, 4}, {0, 3, 4}, {1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4}},
			{{0, 1, 2, 3}, {0, 1, 2, 4}, {0, 1, 3, 4}, {0, 2, 3, 4}, {1, 2, 3, 4}},
			{{0, 1, 2, 3, 4}},
		},

		6: {
			{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 3}, {2, 4}, {2, 5}, {3, 4}, {3, 5}, {4, 5}},
			{{0, 1, 2}, {0, 1, 3}, {0, 1, 4}, {0, 1, 5}, {0, 2, 3}, {0, 2, 4}, {0, 2, 5}, {0, 3, 4}, {0, 3, 5}, {0, 4, 5}, {1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5}},
			{{0, 1, 2, 3}, {0, 1, 2, 4}, {0, 1, 2, 5}, {0, 1, 3, 4}, {0, 1, 3, 5}, {0, 1, 4, 5}, {0, 2, 3, 4}, {0, 2, 3, 5}, {0, 2, 4, 5}, {0, 3, 4, 5}, {1, 2, 3, 4}, {1, 2, 3, 5}, {1, 2, 4, 5}, {1, 3, 4, 5}, {2, 3, 4, 5}},
			{{0, 1, 2, 3, 4}, {0, 1, 2, 3, 5}, {0, 1, 2, 4, 5}, {0, 1, 3, 4, 5}, {0, 2, 3, 4, 5}, {1, 2, 3, 4, 5}},
			{{0, 1, 2, 3, 4, 5}},
		},

		7: {
			{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6},
				{1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6},
				{2, 3}, {2, 4}, {2, 5}, {2, 6},
				{3, 4}, {3, 5}, {3, 6},
				{4, 5}, {4, 6},
				{5, 6}},
			{{0, 1, 2}, {0, 1, 3}, {0, 1, 4}, {0, 1, 5}, {0, 1, 6}, {0, 2, 3}, {0, 2, 4}, {0, 2, 5}, {0, 2, 6}, {0, 3, 4}, {0, 3, 5}, {0, 3, 6}, {0, 4, 5}, {0, 4, 6}, {0, 5, 6},
				{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 2, 6}, {1, 3, 4}, {1, 3, 5}, {1, 3, 6}, {1, 4, 5}, {1, 4, 6}, {1, 5, 6},
				{2, 3, 4}, {2, 3, 5}, {2, 3, 6}, {2, 4, 5}, {2, 4, 6}, {2, 5, 6},
				{3, 4, 5}, {3, 4, 6}, {3, 5, 6},
				{4, 5, 6}},
			{{0, 1, 2, 3}, {0, 1, 2, 4}, {0, 1, 2, 5}, {0, 1, 2, 6}, {0, 1, 3, 4}, {0, 1, 3, 5}, {0, 1, 3, 6}, {0, 1, 4, 5}, {0, 1, 4, 6}, {0, 1, 5, 6},
				{0, 2, 3, 4}, {0, 2, 3, 5}, {0, 2, 3, 6}, {0, 2, 4, 5}, {0, 2, 4, 6}, {0, 2, 5, 6}, {0, 3, 4, 5}, {0, 3, 4, 6}, {0, 3, 5, 6}, {0, 4, 5, 6},
				{1, 2, 3, 4}, {1, 2, 3, 5}, {1, 2, 3, 6}, {1, 2, 4, 5}, {1, 2, 4, 6}, {1, 2, 5, 6}, {1, 3, 4, 5}, {1, 3, 4, 6}, {1, 3, 5, 6}, {1, 4, 5, 6},
				{2, 3, 4, 5}, {2, 3, 4, 6}, {2, 3, 5, 6}, {2, 4, 5, 6},
				{3, 4, 5, 6}},
			{{0, 1, 2, 3, 4}, {0, 1, 2, 3, 5}, {0, 1, 2, 3, 6}, {0, 1, 2, 4, 5}, {0, 1, 2, 4, 6}, {0, 1, 2, 5, 6}, {0, 1, 3, 4, 5}, {0, 1, 3, 4, 6}, {0, 1, 3, 5, 6}, {0, 1, 4, 5, 6},
				{0, 2, 3, 4, 5}, {0, 2, 3, 4, 6}, {0, 2, 3, 5, 6}, {0, 2, 4, 5, 6}, {0, 3, 4, 5, 6},
				{1, 2, 3, 4, 5}, {1, 2, 3, 4, 6}, {1, 2, 3, 5, 6}, {1, 2, 4, 5, 6}, {1, 3, 4, 5, 6},
				{2, 3, 4, 5, 6}},
			{{0, 1, 2, 3, 4, 5}, {0, 1, 2, 3, 4, 6}, {0, 1, 2, 3, 5, 6}, {0, 1, 2, 4, 5, 6}, {0, 1, 3, 4, 5, 6}, {0, 2, 3, 4, 5, 6},
				{1, 2, 3, 4, 5, 6}},
			{{0, 1, 2, 3, 4, 5, 6}},
		},
	}

	combinations, err := AllCombinations(7)
	if err != nil {
		panic(err)
	}
	spew.Dump(combinations[7])

	verifyCombinations(t, combinations, allCombinations)
}

func TestCombinationCount(t *testing.T) {
	allCombinations, err := AllCombinations(10)
	if err != nil {
		panic(err)
	}

	if combinations, ok := allCombinations[2]; ok {
		assert.Equal(t, 1, len(combinations[0]))
	}

	if combinations, ok := allCombinations[3]; ok {
		assert.Equal(t, 3, len(combinations[0]))
		assert.Equal(t, 1, len(combinations[1]))
	}

	if combinations, ok := allCombinations[4]; ok {
		assert.Equal(t, 6, len(combinations[0]))
		assert.Equal(t, 4, len(combinations[1]))
		assert.Equal(t, 1, len(combinations[2]))
	}

	if combinations, ok := allCombinations[5]; ok {
		assert.Equal(t, 10, len(combinations[0]))
		assert.Equal(t, 10, len(combinations[1]))
		assert.Equal(t, 5, len(combinations[2]))
		assert.Equal(t, 1, len(combinations[3]))
	}

	if combinations, ok := allCombinations[6]; ok {
		assert.Equal(t, 15, len(combinations[0]))
		assert.Equal(t, 20, len(combinations[1]))
		assert.Equal(t, 15, len(combinations[2]))
		assert.Equal(t, 6, len(combinations[3]))
		assert.Equal(t, 1, len(combinations[4]))
	}

	if combinations, ok := allCombinations[7]; ok {
		assert.Equal(t, 21, len(combinations[0]))
		assert.Equal(t, 35, len(combinations[1]))
		assert.Equal(t, 35, len(combinations[2]))
		assert.Equal(t, 21, len(combinations[3]))
		assert.Equal(t, 7, len(combinations[4]))
		assert.Equal(t, 1, len(combinations[5]))
	}

	if combinations, ok := allCombinations[8]; ok {
		assert.Equal(t, 28, len(combinations[0]))
		assert.Equal(t, 56, len(combinations[1]))
		assert.Equal(t, 70, len(combinations[2]))
		assert.Equal(t, 56, len(combinations[3]))
		assert.Equal(t, 28, len(combinations[4]))
		assert.Equal(t, 8, len(combinations[5]))
		assert.Equal(t, 1, len(combinations[6]))
	}

	if combinations, ok := allCombinations[9]; ok {
		assert.Equal(t, 36, len(combinations[0]))
		assert.Equal(t, 84, len(combinations[1]))
		assert.Equal(t, 126, len(combinations[2]))
		assert.Equal(t, 126, len(combinations[3]))
		assert.Equal(t, 84, len(combinations[4]))
		assert.Equal(t, 36, len(combinations[5]))
		assert.Equal(t, 9, len(combinations[6]))
		assert.Equal(t, 1, len(combinations[7]))
	}

	if combinations, ok := allCombinations[10]; ok {
		assert.Equal(t, 45, len(combinations[0]))
		assert.Equal(t, 120, len(combinations[1]))
		assert.Equal(t, 210, len(combinations[2]))
		assert.Equal(t, 252, len(combinations[3]))
		assert.Equal(t, 210, len(combinations[4]))
		assert.Equal(t, 120, len(combinations[5]))
		assert.Equal(t, 45, len(combinations[6]))
		assert.Equal(t, 10, len(combinations[7]))
		assert.Equal(t, 1, len(combinations[8]))
	}
}

func verifyCombinations(t *testing.T, result map[int][][][]int, assumption map[int][][][]int) {
	for key, value := range result {
		if assumptionCombinations, ok := assumption[key]; ok {
			for i, combinations := range value {
				for j, combination := range combinations {
					assert.Equal(t, assumptionCombinations[i][j], combination)
				}
			}
		}
	}
}
