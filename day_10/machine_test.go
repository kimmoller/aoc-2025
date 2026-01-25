package main

import (
	"aoc2025/utils"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTurnOn(t *testing.T) {
	data := map[int]string{
		2: "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
		3: "[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
	}

	for key, value := range data {
		machine, err := NewMachine(value)
		if err != nil {
			panic(err)
		}
		verifyMachine(t, value, machine)

		buttonPresses, err := machine.TurnOn()
		if err != nil {
			panic(err)
		}

		assert.Equal(t, key, *buttonPresses)
	}
}

func TestCreateMachine(t *testing.T) {
	data := "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}"
	machine, err := NewMachine(data)
	if err != nil {
		panic(err)
	}

	verifyMachine(t, data, machine)

	assert.Equal(t, []bool{false, true, true, false}, machine.requiredState)
	assert.Equal(t, []int{3}, machine.buttons[0].actions)
	assert.Equal(t, []int{1, 3}, machine.buttons[1].actions)
	assert.Equal(t, []int{2}, machine.buttons[2].actions)
	assert.Equal(t, []int{2, 3}, machine.buttons[3].actions)
	assert.Equal(t, []int{0, 2}, machine.buttons[4].actions)
	assert.Equal(t, []int{0, 1}, machine.buttons[5].actions)
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
}
