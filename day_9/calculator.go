package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type Point struct {
	x int
	y int
}

type Pair struct {
	p1 Point
	p2 Point
}

type Boundry struct {
	min int
	max int
}

func BiggestArea(data []string, withBoundries bool) (*float64, error) {
	pairs, err := toPairs(data, withBoundries)
	if err != nil {
		return nil, err
	}

	biggestArea := float64(0)
	spew.Dump(len(pairs))
	for _, pair := range pairs {
		area := areaBetween(pair.p1, pair.p2)
		if biggestArea < area {
			biggestArea = area
		}
	}
	return &biggestArea, nil
}

func areaBetween(p1 Point, p2 Point) float64 {
	horizontalSide := math.Abs(float64(p1.y-p2.y)) + 1
	verticalSixe := math.Abs(float64(p1.x-p2.x)) + 1

	return horizontalSide * verticalSixe
}

func toPairs(data []string, withBoundries bool) ([]Pair, error) {
	points, err := toPoints(data)
	if err != nil {
		return nil, err
	}

	boundries := toBoundries(points)

	pairs := []Pair{}
	hmm := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			hmm++
			p1 := points[i]
			p2 := points[j]
			if withBoundries && !areaFitsBoundry(p1, p2, boundries) {
				continue
			}
			pair := Pair{p1: p1, p2: p2}
			pairs = append(pairs, pair)
		}
	}
	spew.Dump(hmm)
	return pairs, nil
}

func areaFitsBoundry(p1 Point, p2 Point, boundries map[int]Boundry) bool {
	firstOpposingCorner := Point{x: p2.x, y: p1.y}
	if !pointIsInsideBoundry(firstOpposingCorner, boundries) {
		return false
	}

	secondOpposingCorner := Point{x: p1.x, y: p2.y}
	if !pointIsInsideBoundry(secondOpposingCorner, boundries) {
		return false
	}

	return true
}

func pointIsInsideBoundry(point Point, boundries map[int]Boundry) bool {
	if boundry, ok := boundries[point.y]; ok {
		if point.x < boundry.min || point.x > boundry.max {
			return false
		}
	} else {
		panic(fmt.Errorf("somehow there is no boundry for the Y-coordinate of point %v", point))
	}
	return true
}

func toPoints(data []string) ([]Point, error) {
	points := []Point{}
	for _, item := range data {
		values := strings.Split(item, ",")
		x, err := strconv.Atoi(values[0])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(values[1])
		if err != nil {
			return nil, err
		}
		point := Point{x: x, y: y}
		points = append(points, point)
	}
	return points, nil
}

func toBoundries(points []Point) map[int]Boundry {
	rows := map[int][]int{}
	allPoints := withGapPoints(points)
	for _, point := range allPoints {
		if _, ok := rows[point.y]; ok {
			rows[point.y] = append(rows[point.y], point.x)
		} else {
			rows[point.y] = []int{point.x}
		}
	}
	for key := range rows {
		slices.Sort(rows[key])
	}

	boundries := map[int]Boundry{}
	for y, row := range rows {
		boundries[y] = Boundry{min: row[0], max: row[len(row)-1]}
	}

	return boundries
}

func withGapPoints(points []Point) []Point {
	finalPoints := []Point{}
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		finalPoints = append(finalPoints, p1)
		var p2 Point
		if i == len(points)-1 {
			p2 = points[0]
		} else {
			p2 = points[i+1]
		}

		if p1.x == p2.x {
			if p1.y < p2.y {
				for j := 1; j < p2.y-p1.y; j++ {
					finalPoints = append(finalPoints, Point{x: p1.x, y: p1.y + j})
				}
			} else {
				for j := 1; j < p1.y-p2.y; j++ {
					finalPoints = append(finalPoints, Point{x: p1.x, y: p2.y + j})
				}
			}
		} else {
			if p1.x < p2.x {
				for j := 1; j < p2.x-p1.x; j++ {
					finalPoints = append(finalPoints, Point{x: p1.x + j, y: p1.y})
				}
			} else {
				for j := 1; j < p1.x-p2.x; j++ {
					finalPoints = append(finalPoints, Point{x: p2.x + j, y: p1.y})
				}
			}
		}
	}
	return finalPoints
}
