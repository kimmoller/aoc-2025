package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	SIMPLE  string = "SIMPLE"
	COMPLEX string = "COMPLEX"
)

type Validator struct {
	ids  []int
	rule string
}

func NewValidator(ids []int, rule string) Validator {
	if rule != SIMPLE && rule != COMPLEX {
		panic(fmt.Errorf("Invalid rule %s", rule))
	}
	return Validator{ids, rule}
}

func (v *Validator) InvalidIds() ([]int, error) {
	ids := v.toStrings()
	invalidIds := []string{}
	if v.rule == SIMPLE {
		invalidIds = simpleValidation(ids)
	}
	if v.rule == COMPLEX {
		invalidIds = complexValidation(ids)
	}

	return v.toInts(invalidIds)
}

func simpleValidation(ids []string) []string {
	invalidIds := []string{}
	for _, id := range ids {
		idLength := len(id)
		evenNumber := idLength%2 == 0
		if evenNumber {
			midWay := idLength / 2
			firstHalf := id[:midWay]
			secondHalf := id[midWay:]

			if strings.EqualFold(firstHalf, secondHalf) {
				invalidIds = append(invalidIds, id)
			}
		}
	}
	return invalidIds
}

func complexValidation(ids []string) []string {
	idsWithParts := map[string][][]string{}
	for _, id := range ids {
		parts := splitIntoParts(id)
		idsWithParts[id] = parts
	}

	invalidIds := []string{}
	for id, parts := range idsWithParts {
		valid := isValid(parts)
		if !valid {
			invalidIds = append(invalidIds, id)
		}
	}

	return invalidIds
}

func splitIntoParts(id string) [][]string {
	length := len(id)
	dividers := []int{}
	for i := 1; i < length; i++ {
		if length%i == 0 {
			dividers = append(dividers, i)
		}
	}
	allParts := [][]string{}
	for _, divider := range dividers {
		amountOfParts := length / divider
		startOfSplit := 0
		parts := []string{}
		for i := 0; i < amountOfParts; i++ {
			endOfSplit := startOfSplit + divider
			part := id[startOfSplit:endOfSplit]
			parts = append(parts, part)
			startOfSplit += divider
		}
		allParts = append(allParts, parts)
	}
	return allParts
}

func isValid(allParts [][]string) bool {
	partsValidity := []bool{}
	for _, parts := range allParts {
		partsValidity = append(partsValidity, isValidSetOfParts(parts))
	}

	for _, validPart := range partsValidity {
		if !validPart {
			return false
		}
	}
	return true
}

func isValidSetOfParts(parts []string) bool {
	valid := false
	for i := 1; i < len(parts); i++ {
		if !strings.EqualFold(parts[i-1], parts[i]) {
			valid = true
			break
		}
	}
	return valid
}

func (v *Validator) toStrings() []string {
	idsAsStrings := []string{}
	for _, id := range v.ids {
		idsAsStrings = append(idsAsStrings, strconv.Itoa(id))
	}
	return idsAsStrings
}

func (v *Validator) toInts(ids []string) ([]int, error) {
	idsAsInts := []int{}
	for _, id := range ids {
		intId, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		idsAsInts = append(idsAsInts, intId)
	}
	return idsAsInts, nil
}
