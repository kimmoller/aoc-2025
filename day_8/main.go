package main

import (
	"aoc2025/utils"
)

func Run(path string, connections int) (*int, error) {
	data, err := utils.ReadData(path)
	if err != nil {
		return nil, err
	}

	boxes, err := JunctionBoxes(data)
	if err != nil {
		return nil, err
	}

	circuits := Circuits(connections, boxes)
	sum, err := SumOfThreeLargest(circuits)
	if err != nil {
		return nil, err
	}

	return sum, nil
}
