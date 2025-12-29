package main

import (
	"aoc2025/utils"

	"github.com/davecgh/go-spew/spew"
)

func Day6(path string) (*int, error) {
	data, err := utils.ReadData(path)
	if err != nil {
		return nil, err
	}

	problems, err := ProblemsFromData(data)
	if err != nil {
		return nil, err
	}

	spew.Dump(problems)

	calculator := NewCalculator(problems)

	sum := calculator.SumOfProblems()
	return &sum, nil
}
