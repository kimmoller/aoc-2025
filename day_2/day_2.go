package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ParseData(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	data := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = strings.Split(scanner.Text(), ",")
	}

	return data, nil
}

func SumOfInvalidIds(data []string, withComplex bool) (*int, error) {
	ids, err := getAllIds(data)
	if err != nil {
		return nil, err
	}

	invalidIds, err := findInvalidIds(ids, withComplex)
	if err != nil {
		return nil, err
	}

	sum := 0
	for _, invalidId := range invalidIds {
		sum += invalidId
	}

	return &sum, nil
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

func findInvalidIds(ids []int, withComplex bool) ([]int, error) {
	idsAsStrings := []string{}
	for _, id := range ids {
		idsAsStrings = append(idsAsStrings, strconv.Itoa(id))
	}

	invalidIds := []int{}
	for _, id := range idsAsStrings {
		if withComplex {
			allOneNumber := containsOneNumber(id)
			if allOneNumber {
				intId, err := strconv.Atoi(id)
				if err != nil {
					return nil, err
				}
				invalidIds = append(invalidIds, intId)
				continue
			}
		}

		// TODO: Add other sizes to comparison
		// TODO: Handle odd numbers
		idLength := len(id)
		evenNumber := idLength%2 == 0
		if evenNumber {
			midWay := idLength / 2
			firstHalf := id[:midWay]
			secondHalf := id[midWay:]

			if strings.EqualFold(firstHalf, secondHalf) {
				intId, err := strconv.Atoi(id)
				if err != nil {
					return nil, err
				}
				invalidIds = append(invalidIds, intId)
			}
		} else {

		}
	}

	return invalidIds, nil
}

func containsOneNumber(id string) bool {
	parts := strings.Split(id, "")
	return allEqual(parts)
}

func allEqual(parts []string) bool {
	for i := 1; i < len(parts); i++ {
		if !strings.EqualFold(parts[i-1], parts[i]) {
			return false
		}
	}
	return true
}
