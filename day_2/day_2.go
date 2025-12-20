package main

import (
	"aoc2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day2(path string, rule string) (*int, error) {
	data, err := completeData(path)

	validator := NewValidator(data, rule)
	invalidIds, err := validator.InvalidIds()
	if err != nil {
		return nil, err
	}

	sum := 0
	for _, id := range invalidIds {
		sum += id
	}
	return &sum, nil
}

func completeData(path string) ([]int, error) {
	data, err := utils.ReadData(path)
	if err != nil {
		return nil, err
	}

	parsedData, err := parseData(data)
	if err != nil {
		return nil, err
	}

	ids, err := getAllIds(parsedData)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func parseData(data []string) ([]string, error) {
	if len(data) != 1 {
		return nil, fmt.Errorf("invalid data length %d", len(data))
	}
	splitData := strings.Split(data[0], ",")
	return splitData, nil
}

func getAllIds(data []string) ([]int, error) {
	ids := []int{}
	for _, value := range data {
		boundry := strings.Split(value, "-")
		bottomStr := boundry[0]
		topStr := boundry[1]

		bottom, err := strconv.Atoi(bottomStr)
		if err != nil {
			return nil, err
		}

		top, err := strconv.Atoi(topStr)
		if err != nil {
			return nil, err
		}

		ids = append(ids, bottom)
		amount := top - bottom
		nextId := bottom
		for i := 0; i < amount; i++ {
			nextId = nextId + 1
			ids = append(ids, nextId)
		}
	}
	return ids, nil
}
