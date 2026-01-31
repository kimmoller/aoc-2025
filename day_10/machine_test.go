package main

import (
	"aoc2025/utils"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestTurnOnWihtFiveButtons(t *testing.T) {
	combinationsForFive := [][][]int{
		{
			{0, 1}, {0, 2}, {0, 3}, {0, 4},
			{1, 2}, {1, 3}, {1, 4},
			{2, 3}, {2, 4},
			{3, 4},
		},
		{
			{0, 1, 2}, {0, 1, 3}, {0, 1, 4}, {0, 2, 3}, {0, 2, 4}, {0, 3, 4},
			{1, 2, 3}, {1, 2, 4}, {1, 3, 4},
			{2, 3, 4},
		},
		{
			{0, 1, 2, 3}, {0, 1, 2, 4}, {0, 1, 3, 4}, {0, 2, 3, 4},
			{1, 2, 3, 4},
		},
		{
			{0, 1, 2, 3, 4},
		},
	}

	combinations := map[int][][][]int{
		5: combinationsForFive,
	}

	data := "[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}"

	machine, err := NewMachine(data)
	if err != nil {
		panic(err)
	}
	verifyMachine(t, data, machine)

	buttonPresses, err := machine.TurnOn(combinations)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 3, *buttonPresses)
}

func TestTurnOnWihtSixButtons(t *testing.T) {
	combinationsForSix := [][][]int{
		{
			{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5},
			{1, 2}, {1, 3}, {1, 4}, {1, 5},
			{2, 3}, {2, 4}, {2, 5},
			{3, 4}, {3, 5}},
		{
			{0, 1, 2}, {0, 1, 3}, {0, 1, 4}, {0, 1, 5}, {0, 2, 3}, {0, 2, 4}, {0, 2, 5}, {0, 3, 4}, {0, 3, 5}, {0, 4, 5},
			{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {1, 4, 5},
			{2, 3, 4}, {2, 3, 5}, {2, 4, 5},
			{3, 4, 5},
		},
		{
			{0, 1, 2, 3}, {0, 1, 2, 4}, {0, 1, 3, 4}, {0, 2, 3, 4},
			{1, 2, 3, 4},
		},
		{
			{0, 1, 2, 3, 4}, {0, 1, 2, 3, 5}, {0, 1, 2, 4, 5}, {0, 1, 3, 4, 5}, {0, 2, 3, 4, 5}, {1, 2, 3, 4, 5},
		},
		{
			{0, 1, 2, 3, 4, 5},
		},
	}

	combinations := map[int][][][]int{
		6: combinationsForSix,
	}

	data := "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}"

	machine, err := NewMachine(data)
	if err != nil {
		panic(err)
	}
	verifyMachine(t, data, machine)

	buttonPresses, err := machine.TurnOn(combinations)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 2, *buttonPresses)
}

func TestBrokenMachine(t *testing.T) {
	// 0,1,4,7
	data := "[##..#..#] (0,2,3,4,5,6,7) (5,6) (0,1,2,4) (0,2,3,6) (0,1,2,3,4,7) (2,4,5,6) (1,4,5,6,7) (1,2,4,5,6,7) (0,3,5) {27,38,42,24,48,63,57,40}"
	combinations, err := AllCombinations(9)
	if err != nil {
		panic(err)
	}
	spew.Dump(combinations)

	machine, err := NewMachine(data)
	if err != nil {
		panic(err)
	}

	result, err := machine.TurnOn(combinations)
	if err != nil {
		panic(err)
	}

	spew.Dump(result)
}

func TestCreateAllMachines(t *testing.T) {
	data, err := utils.ReadData("data")
	if err != nil {
		panic(err)
	}

	for _, input := range data {
		machine, err := NewMachine(input)
		if err != nil {
			panic(err)
		}
		verifyMachine(t, input, machine)
	}
}

func verifyMachine(t *testing.T, data string, machine *Machine) {
	inputs := strings.Split(data, " ")
	requiredStateData := strings.Split(inputs[0], "")

	// Simple way to ignore the [] around the data
	amountOfStates := len(requiredStateData) - 2

	assert.Equal(t, amountOfStates, len(machine.initialState))
	assert.Equal(t, amountOfStates, len(machine.requiredState))

	buttonData := inputs[1 : len(inputs)-1]
	assert.Equal(t, len(buttonData), len(machine.buttons))

	joltageData := inputs[len(inputs)-1]
	joltages := strings.Split(joltageData, ",")
	assert.Equal(t, len(joltages), len(machine.joltages))
}
