package main

import (
	"aoc2025/utils"
)

func Day3(path string) (*int, error) {
	data, err := utils.ReadData(path)
	if err != nil {
		return nil, err
	}

	powerSupply, err := NewPowerSupply(data)
	if err != nil {
		return nil, err
	}

	maximumJoltage, err := powerSupply.MaximumJoltage()
	if err != nil {
		return nil, err
	}

	return maximumJoltage, nil
}
