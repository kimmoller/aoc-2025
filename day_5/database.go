package main

import (
	"strconv"
	"strings"
)

type Database struct {
	ranges []string
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

func initDatabase(data []string) ([]string, []int, error) {
	ranges := []string{}
	ids := []int{}
	for _, input := range data {
		if input == "" {
			continue
		}
		isRange := strings.Contains(input, "-")
		if isRange {
			ranges = append(ranges, input)
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
			values := strings.Split(value, "-")
			bottom, err := strconv.Atoi(values[0])
			if err != nil {
				return nil, err
			}
			top, err := strconv.Atoi(values[1])
			if err != nil {
				return nil, err
			}

			if id >= bottom && id <= top {
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
