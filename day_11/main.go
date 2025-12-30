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

	spew.Dump(data)

	sum := 0
	return &sum, nil
}
