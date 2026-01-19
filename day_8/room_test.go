package main

import (
	"aoc2025/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClosestPair(t *testing.T) {
	data, err := utils.ReadData("test_data")
	if err != nil {
		panic(err)
	}

	boxes, err := JunctionBoxes(data)
	if err != nil {
		panic(err)
	}

	distance, box1, box2 := closestPair(0, boxes)

	assert.Equal(t, 316.90219311326956, distance)

	assert.Equal(t, box1.x, 162)
	assert.Equal(t, box1.y, 817)
	assert.Equal(t, box1.z, 812)

	assert.Equal(t, box2.x, 425)
	assert.Equal(t, box2.y, 690)
	assert.Equal(t, box2.z, 689)
}

func TestCircuits(t *testing.T) {
	data, err := utils.ReadData("test_data")
	if err != nil {
		panic(err)
	}

	boxes, err := JunctionBoxes(data)
	if err != nil {
		panic(err)
	}

	circuits := Circuits(10, boxes)
	assert.Equal(t, 5, len(circuits))
}
