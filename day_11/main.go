package main

import (
	"aoc2025/utils"
)

func Run(path string, start string, end string, withMiddleNodes bool) (*int, error) {
	data, err := utils.ReadData(path)
	if err != nil {
		return nil, err
	}

	center := NewCenter()
	err = center.PopulateCenter(data)
	if err != nil {
		return nil, err
	}

	if withMiddleNodes {
		paths, err := center.PathsWithMiddleSteps(start, end)
		if err != nil {
			return nil, err
		}
		return &paths, nil
	}
	paths, err := center.Paths(start, end)
	if err != nil {
		return nil, err
	}

	return &paths, nil
}
