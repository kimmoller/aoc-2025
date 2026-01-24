package main

import (
	"aoc2025/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToPairs(t *testing.T) {
	data := []string{
		"7,1",
		"11,1",
		"11,7",
		"9,7",
	}

	pairs, err := toPairs(data, false)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 6, len(pairs))
}

func TestAreaBetweenPointAndOrigon(t *testing.T) {
	p1 := Point{x: 0, y: 0}
	p2 := Point{x: 4, y: 4}
	area := areaBetween(p1, p2)
	assert.Equal(t, float64(25), area)
}

func TestAreaBetweenTwoPoints(t *testing.T) {
	p1 := Point{x: 1, y: 1}
	p2 := Point{x: 5, y: 5}
	area := areaBetween(p1, p2)
	assert.Equal(t, float64(25), area)
}

func TestAreaBetweenLines(t *testing.T) {
	p1 := Point{x: 0, y: 5}
	p2 := Point{x: 5, y: 5}
	area := areaBetween(p1, p2)
	assert.Equal(t, float64(6), area)

	p1 = Point{x: 5, y: 0}
	p2 = Point{x: 5, y: 5}
	area = areaBetween(p1, p2)
	assert.Equal(t, float64(6), area)
}

func TestToBoundries(t *testing.T) {
	points := []Point{
		{x: 2, y: 2},
		{x: 7, y: 2},
		{x: 7, y: 10},
		{x: 2, y: 10},
	}

	boundries := toBoundries(points)

	assert.Equal(t, 9, len(boundries))

	assert.Equal(t, 2, boundries[2].min)
	assert.Equal(t, 7, boundries[2].max)

	assert.Equal(t, 2, boundries[3].min)
	assert.Equal(t, 7, boundries[3].max)

	assert.Equal(t, 2, boundries[4].min)
	assert.Equal(t, 7, boundries[4].max)

	assert.Equal(t, 2, boundries[5].min)
	assert.Equal(t, 7, boundries[5].max)

	assert.Equal(t, 2, boundries[6].min)
	assert.Equal(t, 7, boundries[6].max)

	assert.Equal(t, 2, boundries[7].min)
	assert.Equal(t, 7, boundries[7].max)

	assert.Equal(t, 2, boundries[8].min)
	assert.Equal(t, 7, boundries[8].max)

	assert.Equal(t, 2, boundries[9].min)
	assert.Equal(t, 7, boundries[9].max)

	assert.Equal(t, 2, boundries[10].min)
	assert.Equal(t, 7, boundries[10].max)
}

func TestGapPoints(t *testing.T) {
	points := []Point{
		{x: 2, y: 2},
		{x: 7, y: 2},
		{x: 7, y: 10},
		{x: 2, y: 10},
	}

	allPoints := withGapPoints(points)
	assert.Equal(t, 26, len(allPoints))
}

func TestToComplexBoundries(t *testing.T) {
	points := []Point{
		{x: 2, y: 2},
		{x: 7, y: 2},
		{x: 10, y: 2},
		{x: 10, y: 5},
		{x: 9, y: 5},
		{x: 9, y: 8},
		{x: 7, y: 8},
		{x: 7, y: 6},
		{x: 4, y: 6},
		{x: 4, y: 4},
		{x: 2, y: 4},
	}

	boundries := toBoundries(points)

	assert.Equal(t, 7, len(boundries))

	assert.Equal(t, 2, boundries[2].min)
	assert.Equal(t, 10, boundries[2].max)

	assert.Equal(t, 2, boundries[4].min)
	assert.Equal(t, 10, boundries[4].max)

	assert.Equal(t, 4, boundries[5].min)
	assert.Equal(t, 10, boundries[5].max)

	assert.Equal(t, 4, boundries[6].min)
	assert.Equal(t, 9, boundries[6].max)

	assert.Equal(t, 7, boundries[8].min)
	assert.Equal(t, 9, boundries[8].max)
}

func TestPointIsInsideBoundries(t *testing.T) {
	points := []Point{
		{x: 2, y: 2},
		{x: 7, y: 2},
		{x: 7, y: 10},
		{x: 2, y: 10},
	}

	boundries := toBoundries(points)

	point := Point{x: 2, y: 2}
	assert.Equal(t, true, pointIsInsideBoundry(point, boundries))

	point = Point{x: 5, y: 2}
	assert.Equal(t, true, pointIsInsideBoundry(point, boundries))
}

func TestAreaIsInsideBoundries(t *testing.T) {
	points := []Point{
		{x: 2, y: 2},
		{x: 7, y: 2},
		{x: 10, y: 2},
		{x: 10, y: 8},
		{x: 6, y: 8},
		{x: 6, y: 4},
		{x: 2, y: 4},
	}

	boundries := toBoundries(points)

	fits := areaFitsBoundry(points[0], points[5], boundries)
	assert.Equal(t, true, fits)

	fits = areaFitsBoundry(points[0], points[3], boundries)
	assert.Equal(t, false, fits)

	fits = areaFitsBoundry(points[2], points[4], boundries)
	assert.Equal(t, true, fits)
}

func TestAreaIsInsideBoundriesWithAlternatingPoints(t *testing.T) {
	points := []Point{
		{x: 2, y: 2},
		{x: 4, y: 2},
		{x: 4, y: 4},
		{x: 6, y: 4},
		{x: 6, y: 6},
		{x: 4, y: 6},
		{x: 4, y: 8},
		{x: 2, y: 8},
		{x: 2, y: 6},
		{x: 0, y: 6},
		{x: 0, y: 4},
		{x: 2, y: 4},
	}

	boundries := toBoundries(points)
	assert.Equal(t, 7, len(boundries))

	fits := areaFitsBoundry(points[0], points[6], boundries)
	assert.Equal(t, true, fits)

	fits = areaFitsBoundry(points[4], points[10], boundries)
	assert.Equal(t, true, fits)

	fits = areaFitsBoundry(points[1], points[8], boundries)
	assert.Equal(t, true, fits)

	fits = areaFitsBoundry(points[0], points[3], boundries)
	assert.Equal(t, false, fits)

	fits = areaFitsBoundry(points[6], points[10], boundries)
	assert.Equal(t, false, fits)
}

func TestWrongFinalPairDoesFitsBoundry(t *testing.T) {
	p1 := Point{x: 18483, y: 86261}
	p2 := Point{x: 81652, y: 13582}

	data, err := utils.ReadData("data")
	if err != nil {
		panic(err)
	}

	points, err := toPoints(data)
	if err != nil {
		panic(err)
	}

	boundries := toBoundries(points)

	isInside := pointIsInsideBoundry(p1, boundries)
	assert.Equal(t, true, isInside)

	isInside = pointIsInsideBoundry(p2, boundries)
	assert.Equal(t, true, isInside)

	fits := areaFitsBoundry(p1, p2, boundries)
	assert.Equal(t, false, fits)
}
