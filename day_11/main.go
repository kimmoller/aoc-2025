package main

import (
	"aoc2025/utils"
)

func Run(path string) (*int, error) {
	data, err := utils.ReadData(path)
	if err != nil {
		return nil, err
	}

	center := NewCenter()
	err = center.PopulateCenter(data)
	if err != nil {
		return nil, err
	}

	paths, err := center.FindAllPaths("you", "out")
	if err != nil {
		return nil, err
	}

	sum := len(paths)
	// sum := 0
	return &sum, nil
}
