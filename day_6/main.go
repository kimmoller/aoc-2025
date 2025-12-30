package main

import (
	"aoc2025/utils"
)

func Run(path string) (*int, error) {
	data, err := utils.ReadData(path)
	if err != nil {
		return nil, err
	}

	problems, err := ProblemsFromData(data)
	if err != nil {
		return nil, err
	}

	calculator := NewCalculator(problems)

	sum := calculator.SumOfProblems()
	return &sum, nil
}
