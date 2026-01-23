package main

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	id int
	x  int
	y  int
	z  int
}

type Pair struct {
	b1       int
	b2       int
	distance float64
}

func SumOfLargest(data []string, connections int) (*int, error) {
	allPoints, boxes, err := points(data)
	if err != nil {
		return nil, err
	}

	storage := NewStorage(boxes)

	allPairs := pairs(allPoints, connections)

	allCircuits, err := circuits(allPairs, storage)
	if err != nil {
		return nil, err
	}

	sizes := []int{}
	for _, circuit := range allCircuits {
		sizes = append(sizes, len(circuit))
	}
	slices.Sort(sizes)
	slices.Reverse(sizes)
	sum := sizes[0] * sizes[1] * sizes[2]
	return &sum, nil
}

func points(data []string) ([]Point, map[int]*JunctionBox, error) {
	points := []Point{}
	boxes := map[int]*JunctionBox{}
	for i, item := range data {
		coordinates := strings.Split(item, ",")
		x, err := strconv.Atoi(coordinates[0])
		if err != nil {
			return nil, nil, err
		}
		y, err := strconv.Atoi(coordinates[1])
		if err != nil {
			return nil, nil, err
		}
		z, err := strconv.Atoi(coordinates[2])
		if err != nil {
			return nil, nil, err
		}
		point := Point{id: i, x: x, y: y, z: z}
		points = append(points, point)
		boxes[i] = &JunctionBox{id: i}
	}
	return points, boxes, nil
}

func pairs(points []Point, numberOfPairs int) []Pair {
	pairs := []Pair{}
	uniqueDistances := map[float64]struct{}{}
	for _, first := range points {
		for _, second := range points {
			if first == second {
				continue
			}

			distance := math.Sqrt(math.Pow(float64((first.x-second.x)), 2) + math.Pow(float64((first.y-second.y)), 2) + math.Pow(float64((first.z-second.z)), 2))
			if _, ok := uniqueDistances[distance]; ok {
				continue
			}

			pair := Pair{b1: first.id, b2: second.id, distance: distance}
			pairs = append(pairs, pair)
			uniqueDistances[distance] = struct{}{}
		}
	}
	slices.SortFunc(pairs, func(a Pair, b Pair) int {
		if a.distance < b.distance {
			return -1
		}
		if a.distance > b.distance {
			return 1
		}
		return 0
	})
	return pairs[:numberOfPairs]
}

func circuits(pairs []Pair, storage *Storage) (map[int][]int, error) {
	allCircuits := map[int][]int{}
	rollingId := 1
	for _, pair := range pairs {
		b1, err := storage.Box(pair.b1)
		if err != nil {
			return nil, err
		}
		b2, err := storage.Box(pair.b2)
		if err != nil {
			return nil, err
		}

		if b1.circuit != 0 && b2.circuit != 0 && (b1.circuit == b2.circuit) {
			// Skipping as both boxes are already in the same circuit
		} else if b1.circuit != 0 && b2.circuit != 0 {
			err := mergeCircuits(b1.circuit, b2.circuit, allCircuits, storage)
			if err != nil {
				return nil, err
			}
		} else if b1.circuit != 0 && b2.circuit == 0 {
			placeIntoCircuit(b2, b1.circuit, allCircuits)
		} else if b1.circuit == 0 && b2.circuit != 0 {
			placeIntoCircuit(b1, b2.circuit, allCircuits)
		} else {
			b1.SetCircuit(rollingId)
			b2.SetCircuit(rollingId)
			allCircuits[rollingId] = []int{b1.id, b2.id}
			rollingId++
		}
	}

	return allCircuits, nil
}

func mergeCircuits(c1 int, c2 int, allCircuits map[int][]int, storage *Storage) error {
	if firstCircuit, ok := allCircuits[c1]; ok {
		if secondCircuit, ok := allCircuits[c2]; ok {
			mergedCircuit := slices.Concat(firstCircuit, secondCircuit)

			allCircuits[c1] = mergedCircuit
			allCircuits[c2] = []int{}

			for _, boxId := range secondCircuit {
				box, err := storage.Box(boxId)
				if err != nil {
					return err
				}
				box.SetCircuit(c1)
			}
		}
	}
	return nil
}

func placeIntoCircuit(box *JunctionBox, circuit int, allCircuits map[int][]int) {
	box.SetCircuit(circuit)
	allCircuits[circuit] = append(allCircuits[circuit], box.id)
}
