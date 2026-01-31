package main

import (
	"fmt"
	"slices"
)

type Combination struct {
	ids []int
}

func AllCombinations(maxCombinations int) (map[int][][][]int, error) {
	allCombinations := map[int][][][]int{}
	items := []int{}
	for i := 0; i < maxCombinations; i++ {
		items = append(items, i)
	}
	for i := 2; i <= maxCombinations; i++ {
		combinations := [][][]int{}
		for j := 2; j <= i; j++ {
			previous := [][]int{}
			if j > 2 {
				previous = combinations[j-3]
			}
			combination, err := Combinations(items[:i], previous, j)
			if err != nil {
				return nil, err
			}
			combinations = append(combinations, combination)
		}
		allCombinations[i] = combinations
	}
	return allCombinations, nil
}

func Combinations(items []int, previousCombinations [][]int, perCombination int) ([][]int, error) {
	if perCombination > len(items) {
		return nil, fmt.Errorf("given combination size %d is bigger than the amount of items, %d", perCombination, len(items))
	}
	if perCombination == len(items) {
		return [][]int{items}, nil
	}
	combinations := [][]int{}
	// Need special handling for two as it only requires a single nested loop
	if perCombination == 2 {
		for i := 0; i < len(items)-1; i++ {
			for j := i + 1; j < len(items); j++ {
				combinations = append(combinations, []int{items[i], items[j]})
			}
		}
		return combinations, nil
	}

	pruned := [][]int{}
	for _, combination := range previousCombinations {
		if combination[len(combination)-1] != items[len(items)-1] {
			pruned = append(pruned, combination)
		}
	}

	newCombinations := [][]int{}
	lastDigit := items[len(items)-1]
	for _, combination := range pruned {
		previousDigit := combination[len(combination)-1]
		if previousDigit == lastDigit-1 {
			clone := slices.Clone(combination)
			newCombination := append(clone, lastDigit)
			newCombinations = append(newCombinations, newCombination)
		} else {
			rollingNumber := 1
			for i := previousDigit; i < lastDigit; i++ {
				clone := slices.Clone(combination)
				newNumber := clone[len(clone)-1] + rollingNumber
				newCombination := append(clone, newNumber)
				newCombinations = append(newCombinations, newCombination)
				rollingNumber++
			}
		}
	}

	return newCombinations, nil
}
