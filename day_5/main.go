package main

import (
	"aoc2025/utils"
)

func Run(path string, withAll bool) (*int, error) {
	data, err := utils.ReadData(path)
	if err != nil {
		return nil, err
	}

	database, err := NewDatabase(data)
	if err != nil {
		return nil, err
	}

	if withAll {
		sum, err := database.AllFreshIds()
		if err != nil {
			return nil, err
		}
		return sum, nil
	}

	freshIds, err := database.FreshIds()
	if err != nil {
		return nil, err
	}

	sum := len(freshIds)
	return &sum, nil
}
