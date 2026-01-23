package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
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
	return invalidIds(ids)
}

func invalidIds(ids []string) []string {
	invalidIds := []string{}
	timeBuildingDividers := int64(0)
	timeValidating := int64(0)
	dividerCache := map[int][]int{}
	for _, id := range ids {
		length := len(id)

		dividers := []int{}
		if cached, ok := dividerCache[length]; ok {
			dividers = cached
		} else {
			for i := 1; i < length; i++ {
				if length%i == 0 {
					dividers = append(dividers, i)
				}
			}
			dividerCache[length] = dividers
		}

		startValidation := time.Now()
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
			if !isValidSetOfParts(parts) {
				invalidIds = append(invalidIds, id)
				break
			}
		}
		done := time.Since(startValidation)
		timeValidating += done.Nanoseconds()
	}
	spew.Dump(fmt.Sprintf("Took %d milliseconds to build dividers", timeBuildingDividers/1000000))
	spew.Dump(fmt.Sprintf("Took %d milliseconds validating", timeValidating/1000000))
	return invalidIds
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
