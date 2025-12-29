package main

import (
	"strings"
)

type Warehouse struct {
	width     int
	inventory map[int][]bool
}

func NewWarehouse() *Warehouse {
	return &Warehouse{}
}

func (w *Warehouse) Fill(data []string) {
	w.width = len(data[0])
	inventory := map[int][]bool{}
	for i, line := range data {
		row := []bool{}
		values := strings.Split(line, "")
		for _, value := range values {
			if value == "@" {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		inventory[i] = row
	}
	w.inventory = inventory
}

func (w *Warehouse) NumberOfAccessibleRolls() int {
	sum := 0
	for index, row := range w.inventory {
		for spot, roll := range row {
			if !roll {
				continue
			}

			adjacentSpots := w.adjacentSpots(spot, index)
			adjacentRolls := 0
			for _, spot := range adjacentSpots {
				if spot {
					adjacentRolls++
				}
			}

			if adjacentRolls < 4 {
				sum++
			}
		}
	}

	return sum
}

func (w *Warehouse) adjacentSpots(spot int, row int) []bool {
	adjacentSpots := []bool{}

	// Get spots for the row above
	adjacentSpots = append(adjacentSpots, partialAdjacentSpots(row-1, spot, w.width, w.inventory, false)...)
	// Get spots for the same row
	adjacentSpots = append(adjacentSpots, partialAdjacentSpots(row, spot, w.width, w.inventory, true)...)
	// Get spots for the row below
	adjacentSpots = append(adjacentSpots, partialAdjacentSpots(row+1, spot, w.width, w.inventory, false)...)

	return adjacentSpots
}

func partialAdjacentSpots(rowIndex int, spot int, warehouseWidth int, inventory map[int][]bool, isSameRow bool) []bool {
	spots := []bool{}
	if row, ok := inventory[rowIndex]; ok {
		locations := []int{spot - 1, spot, spot + 1}
		if isSameRow {
			locations = []int{spot - 1, spot + 1}
		}
		for _, location := range locations {
			if location >= 0 && location < warehouseWidth {
				roll := row[location]
				spots = append(spots, roll)
			}
		}
	}
	return spots
}
