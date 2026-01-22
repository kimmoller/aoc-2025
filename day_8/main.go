package main

import (
	"aoc2025/utils"
)

func Run(path string, connections int) (*int, error) {
	data, err := utils.ReadData(path)
	if err != nil {
		return nil, err
	}

	sum, err := SumOfLargest(data, connections)

	return sum, nil
}
