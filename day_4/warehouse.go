package main

import (
	"strings"
)

type Location struct {
	x      int
	y      int
	isRoll bool
}

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

func (w *Warehouse) RemoveRolls(locations []Location) {
	for _, location := range locations {
		row := w.inventory[location.y]
		row[location.x] = false
	}
}

func (w *Warehouse) RecursiveAccessibleRollLocations(currentAccessibleRolls []Location) []Location {
	accessibleRolls := currentAccessibleRolls
	newAccessibleRolls := w.AccessibleRollLocations()

	if len(newAccessibleRolls) > 0 {
		w.RemoveRolls(newAccessibleRolls)
		accessibleRolls = append(accessibleRolls, newAccessibleRolls...)
		newRollCount := w.RecursiveAccessibleRollLocations(currentAccessibleRolls)
		accessibleRolls = append(accessibleRolls, newRollCount...)
	}

	return accessibleRolls
}

func (w *Warehouse) AccessibleRollLocations() []Location {
	accessibleRolls := []Location{}
	for index, row := range w.inventory {
		for spot, roll := range row {
			if !roll {
				continue
			}

			adjacentLocations := w.adjacentLocations(spot, index)
			adjacentRolls := 0
			for _, location := range adjacentLocations {
				if location.isRoll {
					adjacentRolls++
				}
			}

			if adjacentRolls < 4 {
				location := Location{x: spot, y: index, isRoll: true}
				accessibleRolls = append(accessibleRolls, location)
			}
		}
	}

	return accessibleRolls
}

func (w *Warehouse) adjacentLocations(spot int, row int) []Location {
	adjacentLocations := []Location{}

	// Get locations for the row above
	adjacentLocations = append(adjacentLocations, partialAdjacentLocations(row-1, spot, w.width, w.inventory, false)...)
	// Get locations for the same row
	adjacentLocations = append(adjacentLocations, partialAdjacentLocations(row, spot, w.width, w.inventory, true)...)
	// Get locations for the row below
	adjacentLocations = append(adjacentLocations, partialAdjacentLocations(row+1, spot, w.width, w.inventory, false)...)

	return adjacentLocations
}

func partialAdjacentLocations(rowIndex int, spot int, warehouseWidth int, inventory map[int][]bool, isSameRow bool) []Location {
	locations := []Location{}
	if row, ok := inventory[rowIndex]; ok {
		spots := []int{spot - 1, spot, spot + 1}
		if isSameRow {
			spots = []int{spot - 1, spot + 1}
		}
		for _, spot := range spots {
			if spot >= 0 && spot < warehouseWidth {
				roll := row[spot]
				location := Location{x: spot, y: rowIndex, isRoll: roll}
				locations = append(locations, location)
			}
		}
	}
	return locations
}
