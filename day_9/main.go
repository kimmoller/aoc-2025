package main

import (
	"aoc2025/utils"
)

func Run(path string, withBoundries bool) (*float64, error) {
	data, err := utils.ReadData(path)
	if err != nil {
		return nil, err
	}

	area, err := BiggestArea(data, withBoundries)
	if err != nil {
		return nil, err
	}

	return area, nil
}
