package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistanceCalculation(t *testing.T) {
	p1 := Point{id: 0, x: 162, y: 817, z: 812}
	p2 := Point{id: 1, x: 57, y: 618, z: 57}

	distance := distanceBetweenPoints(p1, p2)
	assert.Equal(t, float64(787.814064357828), distance)
}

func TestPoints(t *testing.T) {
	data := []string{
		"162,817,812",
		"57,618,57",
		"906,360,560",
		"592,479,940",
		"352,342,300",
		"466,668,158",
		"542,29,236",
		"431,825,988",
		"739,650,466",
		"52,470,668",
		"216,146,977",
		"819,987,18",
		"117,168,530",
		"805,96,715",
		"346,949,466",
		"970,615,88",
		"941,993,340",
		"862,61,35",
		"984,92,344",
		"425,690,689",
	}

	allPoints, boxes, err := points(data)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 20, len(allPoints))
	assert.Equal(t, 20, len(boxes))

	assert.Equal(t, 162, allPoints[0].x)
	assert.Equal(t, 817, allPoints[0].y)
	assert.Equal(t, 812, allPoints[0].z)

	for _, point := range allPoints {
		if _, ok := boxes[point.id]; !ok {
			panic(fmt.Errorf("Point %d does not exist in boxes", point.id))
		}
	}
}

func TestPairs(t *testing.T) {
	allPoints := []Point{
		{id: 0, x: 162, y: 817, z: 812},
		{id: 1, x: 57, y: 618, z: 57},
		{id: 2, x: 906, y: 360, z: 560},
		{id: 3, x: 592, y: 479, z: 940},
	}

	allPairs := pairs(allPoints)

	distanceFrom0To1 := distanceBetweenPoints(allPoints[0], allPoints[1])
	distanceFrom0To2 := distanceBetweenPoints(allPoints[0], allPoints[2])
	distanceFrom0To3 := distanceBetweenPoints(allPoints[0], allPoints[3])

	distanceFrom1To2 := distanceBetweenPoints(allPoints[1], allPoints[2])
	distanceFrom1To3 := distanceBetweenPoints(allPoints[1], allPoints[3])

	distanceFrom2To3 := distanceBetweenPoints(allPoints[2], allPoints[3])

	assert.Equal(t, 6, len(allPairs))

	assert.Equal(t, distanceFrom2To3, allPairs[0].distance)
	assert.Equal(t, distanceFrom0To3, allPairs[1].distance)
	assert.Equal(t, distanceFrom0To1, allPairs[2].distance)
	assert.Equal(t, distanceFrom0To2, allPairs[3].distance)
	assert.Equal(t, distanceFrom1To2, allPairs[4].distance)
	assert.Equal(t, distanceFrom1To3, allPairs[5].distance)
}

func TestMergeCircuits(t *testing.T) {
	boxes := map[int]*JunctionBox{
		0: {id: 0, circuit: 1},
		1: {id: 1, circuit: 1},
		2: {id: 2, circuit: 1},
		3: {id: 3, circuit: 2},
		4: {id: 4, circuit: 2},
	}
	storage := NewStorage(boxes)
	circuits := map[int][]int{
		1: {0, 1, 2},
		2: {3, 4},
	}

	err := mergeCircuits(1, 2, circuits, storage)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1, len(circuits))
	assert.Equal(t, 5, len(circuits[1]))
	for _, box := range boxes {
		assert.Equal(t, 1, box.circuit)
	}
}

func TestLimitedCircuits(t *testing.T) {
	data := []string{
		"162,817,812",
		"57,618,57",
		"906,360,560",
		"592,479,940",
		"352,342,300",
		"466,668,158",
		"542,29,236",
		"431,825,988",
		"739,650,466",
		"52,470,668",
		"216,146,977",
		"819,987,18",
		"117,168,530",
		"805,96,715",
		"346,949,466",
		"970,615,88",
		"941,993,340",
		"862,61,35",
		"984,92,344",
		"425,690,689",
	}

	allPoints, boxes, err := points(data)
	if err != nil {
		panic(err)
	}

	storage := NewStorage(boxes)
	allPairs := pairs(allPoints)
	smallPairs := limitedPairs(allPairs, 4)

	allCircuits, _, err := circuits(smallPairs, storage)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 2, len(allCircuits))

	firstCircuit := allCircuits[1]
	assert.Equal(t, 0, firstCircuit[0])
	assert.Equal(t, 19, firstCircuit[1])
	assert.Equal(t, 7, firstCircuit[2])

	secondCircuit := allCircuits[2]
	assert.Equal(t, 2, secondCircuit[0])
	assert.Equal(t, 13, secondCircuit[1])
}

func TestLimitlessCircuits(t *testing.T) {
	data := []string{
		"162,817,812",
		"57,618,57",
		"906,360,560",
		"592,479,940",
		"352,342,300",
		"466,668,158",
		"542,29,236",
		"431,825,988",
		"739,650,466",
		"52,470,668",
		"216,146,977",
		"819,987,18",
		"117,168,530",
		"805,96,715",
		"346,949,466",
		"970,615,88",
		"941,993,340",
		"862,61,35",
		"984,92,344",
		"425,690,689",
	}

	allPoints, boxes, err := points(data)
	if err != nil {
		panic(err)
	}

	storage := NewStorage(boxes)
	allPairs := pairs(allPoints)

	allCircuits, lastPair, err := circuits(allPairs, storage)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1, len(allCircuits))

	assert.Equal(t, 10, lastPair.b1)
	assert.Equal(t, 12, lastPair.b2)
}
