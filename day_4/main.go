package main

import (
	"aoc2025/utils"
)

func Run(path string, recursive bool) (*int, error) {
	data, err := utils.ReadData(path)
	if err != nil {
		return nil, err
	}

	warehouse := NewWarehouse()
	warehouse.Fill(data)

	if recursive {
		locations := warehouse.RecursiveAccessibleRollLocations([]Location{})
		sum := len(locations)
		return &sum, nil
	}

	locations := warehouse.AccessibleRollLocations()
	sum := len(locations)
	return &sum, nil
}
