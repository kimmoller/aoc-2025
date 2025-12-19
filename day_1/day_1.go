package main

import (
	"strconv"
)

func CalculatePassword(data []string) (*int, error) {
	currentPoint := 50
	atZero := 0
	for _, item := range data {
		direction := item[:1]

		strValue := item[1:]

		// Input values can be larger than 99
		trimmedValue := strValue
		if len(strValue) > 2 {
			trimmedValue = strValue[1:3]
		}

		value, err := strconv.Atoi(trimmedValue)
		if err != nil {
			return nil, err
		}

		newPoint, hitZero, _ := calculateCurrentPoint(currentPoint, value, direction)
		currentPoint = newPoint

		if hitZero {
			atZero++
		}
	}
	return &atZero, nil
}

func CalculatePasswordWihtRotations(data []string) (*int, error) {
	currentPoint := 50
	atZero := 0
	for _, item := range data {
		direction := item[:1]

		strValue := item[1:]

		// Input values can be larger than 99
		trimmedValue := strValue
		if len(strValue) > 2 {
			trimmedValue = strValue[1:3]
			firstDigitStr := strValue[:1]
			firstDigit, err := strconv.Atoi(firstDigitStr)
			if err != nil {
				return nil, err
			}
			atZero += firstDigit
		}

		value, err := strconv.Atoi(trimmedValue)
		if err != nil {
			return nil, err
		}

		wasAtZero := currentPoint == 0
		newPoint, hitZero, passedZero := calculateCurrentPoint(currentPoint, value, direction)
		currentPoint = newPoint

		if (hitZero || passedZero) && !wasAtZero {
			atZero++
		}
	}
	return &atZero, nil
}

func calculateCurrentPoint(currentPoint int, value int, direction string) (int, bool, bool) {
	if direction == "L" {
		currentPoint = currentPoint - value
	} else {
		currentPoint = currentPoint + value
	}

	hitZero := false
	passedZero := false
	if currentPoint == 100 {
		currentPoint = 0
	}

	if currentPoint < 0 {
		passedZero = true
		currentPoint = currentPoint + 100
	}

	if currentPoint > 100 {
		passedZero = true
		currentPoint = currentPoint - 100
	}

	if currentPoint == 0 {
		hitZero = true
	}

	return currentPoint, hitZero, passedZero
}
