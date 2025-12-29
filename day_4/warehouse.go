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

	adjacentSpots = append(adjacentSpots, spotsAbove(row, spot, w.width, w.inventory)...)
	adjacentSpots = append(adjacentSpots, spotsNextTo(row, spot, w.width, w.inventory)...)
	adjacentSpots = append(adjacentSpots, spotsBelow(row, spot, w.width, w.inventory)...)

	return adjacentSpots
}

func spotsAbove(row int, spot int, warehouseWidth int, inventory map[int][]bool) []bool {
	spots := []bool{}
	if row, ok := inventory[row-1]; ok {

		locations := []int{spot - 1, spot, spot + 1}
		for _, location := range locations {
			if location >= 0 && location < warehouseWidth {
				roll := row[location]
				spots = append(spots, roll)
			}
		}

	}
	return spots
}

func spotsNextTo(row int, spot int, warehouseWidth int, inventory map[int][]bool) []bool {
	spots := []bool{}
	if row, ok := inventory[row]; ok {

		locations := []int{spot - 1, spot + 1}
		for _, location := range locations {
			if location >= 0 && location < warehouseWidth {
				roll := row[location]
				spots = append(spots, roll)
			}
		}

	}
	return spots
}

func spotsBelow(row int, spot int, warehouseWidth int, inventory map[int][]bool) []bool {
	spots := []bool{}
	if row, ok := inventory[row+1]; ok {

		locations := []int{spot - 1, spot, spot + 1}
		for _, location := range locations {
			if location >= 0 && location < warehouseWidth {
				roll := row[location]
				spots = append(spots, roll)
			}
		}

	}
	return spots
}
