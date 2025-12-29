package main

import (
	"aoc2025/utils"
)

func Day4(path string, amountToActivate int) (*int, error) {
	data, err := utils.ReadData(path)
	if err != nil {
		return nil, err
	}

	warehouse := NewWarehouse()
	warehouse.Fill(data)

	sum := warehouse.NumberOfAccessibleRolls()
	return &sum, nil
}
