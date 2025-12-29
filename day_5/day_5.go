package main

import (
	"aoc2025/utils"
)

func Day5(path string) (*int, error) {
	data, err := utils.ReadData(path)
	if err != nil {
		return nil, err
	}

	database, err := NewDatabase(data)
	if err != nil {
		return nil, err
	}

	freshIds, err := database.FreshIds()
	if err != nil {
		return nil, err
	}

	sum := len(freshIds)
	return &sum, nil
}

func Day5_part2(path string) (*int, error) {
	data, err := utils.ReadData(path)
	if err != nil {
		return nil, err
	}

	database, err := NewDatabase(data)
	if err != nil {
		return nil, err
	}

	sum, err := database.AllFreshIds()
	if err != nil {
		return nil, err
	}

	// sum := len(ids)

	return sum, nil
}
