package main

import (
	"aoc2025/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnownPaths(t *testing.T) {
	data, err := utils.ReadData("test_data_2")
	if err != nil {
		panic(err)
	}

	center := NewCenter()
	err = center.PopulateCenter(data)

	startingServer, err := center.Server("svr")
	if err != nil {
		panic(err)
	}
	paths, err := center.FindPath([]string{}, startingServer.name, "out")
	assert.Equal(t, 8, len(paths))
	assert.Equal(t, 8, len(startingServer.paths))
}
