package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestButtonCombinationsWith2(t *testing.T) {
	buttons := []Button{
		{actions: []int{3}},
		{actions: []int{1, 3}},
		{actions: []int{2}},
		{actions: []int{2, 3}},
		{actions: []int{0, 2}},
		{actions: []int{0, 1}},
	}

	combinations, err := Combinations(2, buttons)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 15, len(combinations))
}

func TestButtonCombinationsWith3(t *testing.T) {
	buttons := []Button{
		{actions: []int{3}},
		{actions: []int{1, 3}},
		{actions: []int{2}},
		{actions: []int{2, 3}},
		{actions: []int{0, 2}},
		{actions: []int{0, 1}},
	}

	combinations, err := Combinations(3, buttons)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 20, len(combinations))
}

func TestButtonCombinationsWith4(t *testing.T) {
	buttons := []Button{
		{actions: []int{3}},
		{actions: []int{1, 3}},
		{actions: []int{2}},
		{actions: []int{2, 3}},
		{actions: []int{0, 2}},
		{actions: []int{0, 1}},
	}

	combinations, err := Combinations(4, buttons)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 15, len(combinations))
}

func TestButtonCombinationsWith5(t *testing.T) {
	buttons := []Button{
		{actions: []int{3}},
		{actions: []int{1, 3}},
		{actions: []int{2}},
		{actions: []int{2, 3}},
		{actions: []int{0, 2}},
		{actions: []int{0, 1}},
	}

	combinations, err := Combinations(5, buttons)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 6, len(combinations))
}

func TestButtonCombinationsWith6(t *testing.T) {
	buttons := []Button{
		{actions: []int{3}},
		{actions: []int{1, 3}},
		{actions: []int{2}},
		{actions: []int{2, 3}},
		{actions: []int{0, 2}},
		{actions: []int{0, 1}},
	}

	combinations, err := Combinations(6, buttons)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 1, len(combinations))
}

func TestButtonCombinationsErrorWithTooManyCombinations(t *testing.T) {
	data := "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}"
	machine, err := NewMachine(data)
	if err != nil {
		panic(err)
	}
	verifyMachine(t, data, machine)

	_, err = machine.buttonCombinations(7)
	assert.Error(t, err)
}

// func TestButtonCombinations(t *testing.T) {
// 	data := "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}"
// 	machine, err := NewMachine(data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	verifyMachine(t, data, machine)

// 	combinations := machine.ButtonCombinations(3)
// 	assert.Equal(t, 20, len(combinations))

// 	combinations = machine.ButtonCombinations(4)
// 	assert.Equal(t, 15, len(combinations))

// 	combinations = machine.ButtonCombinations(5)
// 	assert.Equal(t, 6, len(combinations))

// 	combinations = machine.ButtonCombinations(6)
// 	assert.Equal(t, 1, len(combinations))
// }
