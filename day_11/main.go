package main

import (
	"aoc2025/utils"
)

func Run(path string, start string, end string, strict bool) (*int, error) {
	data, err := utils.ReadData(path)
	if err != nil {
		return nil, err
	}

	center := NewCenter()
	err = center.PopulateCenter(data)
	if err != nil {
		return nil, err
	}

	paths, err := center.FindAllPaths(start, end, strict)
	if err != nil {
		return nil, err
	}

	sum := len(paths)
	return &sum, nil
}

func RunSumBased(path string, start, end string, strict bool) (*int, error) {
	data, err := utils.ReadData(path)
	if err != nil {
		return nil, err
	}

	center := NewCenter()
	err = center.PopulateCenter(data)
	if err != nil {
		return nil, err
	}

	sum, err := center.AmountOfValidPaths(start, end, strict)
	if err != nil {
		return nil, err
	}

	return sum, nil
}
