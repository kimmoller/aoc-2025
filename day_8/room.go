package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type JunctionBox struct {
	id      int
	x       int
	y       int
	z       int
	circuit int
}

func JunctionBoxes(data []string) (map[int]JunctionBox, error) {
	boxes := map[int]JunctionBox{}
	for i, point := range data {
		points := strings.Split(point, ",")
		x, err := strconv.Atoi(points[0])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(points[1])
		if err != nil {
			return nil, err
		}
		z, err := strconv.Atoi(points[2])
		if err != nil {
			return nil, err
		}
		box := JunctionBox{id: i, x: x, y: y, z: z}
		boxes[i] = box
	}
	return boxes, nil
}

func SumOfThreeLargest(circuits map[int][]JunctionBox) (*int, error) {
	sizes := []int{}
	for _, circuit := range circuits {
		sizes = append(sizes, len(circuit))
	}
	slices.Sort(sizes)
	slices.Reverse(sizes)
	if len(sizes) < 3 {
		return nil, fmt.Errorf("Not enough circuits")
	}

	sum := sizes[0] * sizes[1] * sizes[2]
	return &sum, nil
}

func Circuits(connections int, boxes map[int]JunctionBox) map[int][]JunctionBox {
	lastDistance := float64(0)
	rollingId := 1
	circuits := map[int][]JunctionBox{}
	for i := 0; i < connections; i++ {
		distance, box1, box2 := closestPair(lastDistance, boxes)
		lastDistance = distance

		// They are in the same circuit, nothing happens
		if box1.circuit != 0 && box2.circuit != 0 && (box1.circuit == box2.circuit) {
			continue
		}

		// Both are in different circuits so the circuits are merged
		if box1.circuit != 0 && box2.circuit != 0 {
			if _, ok := circuits[box1.circuit]; ok {
				if _, ok := circuits[box2.circuit]; ok {
					secondCircuit := slices.Clone(circuits[box2.circuit])
					for j := 0; j < len(secondCircuit); j++ {
						secondCircuit[j].circuit = box1.circuit
					}
					circuits[box1.circuit] = append(circuits[box1.circuit], secondCircuit...)
					circuits[box2.circuit] = []JunctionBox{}
				}
			}
			continue
		}

		// Place Box 2 into Box 1 circuit
		if box1.circuit != 0 && box2.circuit == 0 {
			if box, ok := boxes[box2.id]; ok {
				box.circuit = box1.circuit
				boxes[box.id] = box
				if _, ok := circuits[box1.circuit]; ok {
					circuits[box1.circuit] = append(circuits[box1.circuit], box)
				}
			}
			continue
		}

		// Place Box 1 into Box 2 circuit
		if box1.circuit == 0 && box2.circuit != 0 {
			if box, ok := boxes[box1.id]; ok {
				box.circuit = box2.circuit
				boxes[box.id] = box
				if _, ok := circuits[box2.circuit]; ok {
					circuits[box2.circuit] = append(circuits[box2.circuit], box)
				}
			}
			continue
		}

		circuit := []JunctionBox{}
		if box, ok := boxes[box1.id]; ok {
			box.circuit = rollingId
			boxes[box.id] = box
			circuit = append(circuit, box)
		}
		if box, ok := boxes[box2.id]; ok {
			box.circuit = rollingId
			boxes[box.id] = box
			circuit = append(circuit, box)
		}
		circuits[rollingId] = circuit
		rollingId++
	}
	return circuits
}

func closestPair(lastDistance float64, boxes map[int]JunctionBox) (float64, JunctionBox, JunctionBox) {
	shortestDistance := float64(0)
	var box1 JunctionBox
	var box2 JunctionBox
	for _, first := range boxes {
		for _, second := range boxes {
			if first == second {
				continue
			}
			distance := math.Sqrt(math.Pow(float64((first.x-second.x)), 2) + math.Pow(float64((first.y-second.y)), 2) + math.Pow(float64((first.z-second.z)), 2))
			if (shortestDistance == 0 || shortestDistance > distance) && distance > lastDistance {
				shortestDistance = distance
				box1 = first
				box2 = second
			}
		}
	}
	return shortestDistance, box1, box2
}
