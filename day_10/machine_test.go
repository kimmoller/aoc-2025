package main

import (
	"aoc2025/utils"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	assert.Equal(t, amountOfStates, len(machine.currentState))
	assert.Equal(t, amountOfStates, len(machine.requiredState))

	buttonData := inputs[1 : len(inputs)-1]
	assert.Equal(t, len(buttonData), len(machine.buttons))
}
