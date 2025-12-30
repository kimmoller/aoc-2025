package main

import (
	"aoc2025/utils"
)

func Run(path string) (*int, error) {
	data, err := utils.ReadData(path)
	if err != nil {
		return nil, err
	}

	teleporter := NewTeleporter(data)
	teleporter.Start()

	sum := len(teleporter.splitPositions)
	return &sum, nil
}
