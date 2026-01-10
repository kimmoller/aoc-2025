package main

import (
	"aoc2025/utils"

	"github.com/davecgh/go-spew/spew"
)

func Run(path string) (*int, error) {
	data, err := utils.ReadData(path)
	if err != nil {
		return nil, err
	}

	teleporter := NewTeleporter(data)
	teleporter.Start()

	spew.Dump(len(teleporter.beams))
	sum := len(teleporter.splitPositions)
	return &sum, nil
}
