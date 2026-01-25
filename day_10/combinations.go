package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type Combination struct {
	buttons []Button
}

func Combinations(number int, buttons []Button) ([]Combination, error) {
	spew.Dump(fmt.Sprintf("Get combinations for size %d", number))
	finalCombinations := [][]Combination{}
	if number == 2 || number == len(buttons) {
		combinations, err := calculateCombinations(number, [][]Combination{}, buttons)
		if err != nil {
			return nil, err
		}
		finalCombinations = append(finalCombinations, combinations...)
	} else {
		baseCombinations := [][]Combination{}
		for i := 2; i < number; i++ {
			newBase, err := calculateCombinations(i, baseCombinations, buttons)
			if err != nil {
				return nil, err
			}
			pruned := pruneCombinations(newBase)
			baseCombinations = pruned
		}

		combinations, err := calculateCombinations(number, baseCombinations, buttons)
		if err != nil {
			return nil, err
		}
		finalCombinations = append(finalCombinations, combinations...)
	}

	flatCombinations := []Combination{}
	for _, combination := range finalCombinations {
		flatCombinations = append(flatCombinations, combination...)
	}
	// spew.Dump(flatCombinations)
	return flatCombinations, nil
}

func calculateCombinations(number int, baseCombinations [][]Combination, buttons []Button) ([][]Combination, error) {
	is4 := false
	if number == 4 {
		is4 = true
	}
	if number > len(buttons) {
		return nil, fmt.Errorf("given combination size %d is bigger than the amount of buttons, %d", number, len(buttons))
	}
	if number == len(buttons) {
		return [][]Combination{{{buttons: buttons}}}, nil
	}
	combinations := [][]Combination{}
	// Need special handling for two as it only requires a single nested loop
	if number == 2 {
		for i := 0; i < len(buttons)-1; i++ {
			newCombinations := []Combination{}
			for j := i + 1; j < len(buttons); j++ {
				newCombinations = append(newCombinations, Combination{buttons: []Button{buttons[i], buttons[j]}})
			}
			combinations = append(combinations, newCombinations)
		}
		return combinations, nil
	}

	for i := 0; i < len(baseCombinations); i++ {
		if is4 {
			spew.Dump(fmt.Sprintf("First loop index %d", i))
		}
		subSet := baseCombinations[i]
		for j := 0; j < len(subSet); j++ {
			if is4 {
				spew.Dump(fmt.Sprintf("Second loop index %d", j))
			}
			newCombination := []Combination{}
			baseCombination := subSet[j]
			if is4 {
				spew.Dump(fmt.Sprintf("Base combination %v", baseCombination))
			}
			startingIndex := 0
			// for i, button := range buttons {
			// 	if button == baseCombination.buttons[0] {

			// 	}
			// }
			// startingIndex :=  len(baseCombination.buttons) + j + i
			for k := startingIndex; k < len(buttons); k++ {
				if is4 {
					spew.Dump(fmt.Sprintf("Starting index %d", k))
				}
				combinationButtons := append(baseCombination.buttons, buttons[k])
				newCombination = append(newCombination, Combination{buttons: combinationButtons})
				if is4 {
					spew.Dump(fmt.Sprintf("Appended combination %v", combinationButtons))
				}
			}
			if is4 {
				spew.Dump(fmt.Sprintf("Append new combination %v", newCombination))
			}
			combinations = append(combinations, newCombination)
		}
	}

	return combinations, nil
}

func pruneCombinations(combinations [][]Combination) [][]Combination {
	// Prune out all lists with only one combination
	validBases := [][]Combination{}
	for _, combination := range combinations {
		if len(combination) > 1 {
			validBases = append(validBases, combination)
		}
	}

	baseCombinations := [][]Combination{}
	for _, combination := range validBases {
		baseCombinations = append(baseCombinations, combination[:len(combination)-1])
	}

	return baseCombinations
}
