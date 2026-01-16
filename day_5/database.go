package main

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

type Database struct {
	ranges []Range
	ids    []int
}

func NewDatabase(data []string) (*Database, error) {
	ranges, ids, err := initDatabase(data)
	if err != nil {
		return nil, err
	}
	database := Database{ranges: ranges, ids: ids}

	return &database, nil
}

func initDatabase(data []string) ([]Range, []int, error) {
	ranges := []Range{}
	ids := []int{}
	for _, input := range data {
		if input == "" {
			continue
		}
		isRange := strings.Contains(input, "-")
		if isRange {
			values := strings.Split(input, "-")
			start, err := strconv.Atoi(values[0])
			if err != nil {
				return nil, nil, err
			}
			end, err := strconv.Atoi(values[1])
			if err != nil {
				return nil, nil, err
			}
			numRange := Range{start: start, end: end}
			ranges = append(ranges, numRange)
		} else {
			id, err := strconv.Atoi(input)
			if err != nil {
				return nil, nil, err
			}
			ids = append(ids, id)
		}
	}
	return ranges, ids, nil
}

func (d *Database) FreshIds() ([]int, error) {
	freshIds := []int{}
	for _, id := range d.ids {
		isFresh := false
		for _, value := range d.ranges {
			if id >= value.start && id <= value.end {
				isFresh = true
				break
			}
		}

		if isFresh {
			freshIds = append(freshIds, id)
		}
	}
	return freshIds, nil
}

func (d *Database) AllFreshIds() (*int, error) {
	freshIds := 0
	slices.SortFunc(d.ranges, func(a, b Range) int {
		if a.start == b.start {
			return cmp.Compare(a.end, b.end)
		}
		return cmp.Compare(a.start, b.start)
	})

	mergedRanges := []Range{}
	current := d.ranges[0]
	for i := 1; i < len(d.ranges); i++ {
		next := d.ranges[i]
		if next.start > current.end {
			mergedRanges = append(mergedRanges, current)
			current = next
			continue
		}

		if next.start <= current.end && next.end <= current.end {
			continue
		}

		current.end = next.end
	}
	mergedRanges = append(mergedRanges, current)

	for _, value := range mergedRanges {
		freshIds += value.end - value.start + 1
	}

	return &freshIds, nil
}
