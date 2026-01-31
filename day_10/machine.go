package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Machine struct {
	initialState  []bool
	requiredState []bool
	buttons       map[int][]int
}

func NewMachine(data string) (*Machine, error) {
	return toMachine(data)
}

func (m *Machine) TurnOn(combinationsPerNumberOfItems map[int][][][]int) (*int, error) {
	// Verify if the machine can be turned on with one button press
	for _, actions := range m.buttons {
		currentState := slices.Clone(m.initialState)
		for _, action := range actions {
			currentState[action] = toggleState(currentState[action])
		}
		if isOn(currentState, m.requiredState) {
			result := 1
			return &result, nil
		}
	}

	if allCombinations, ok := combinationsPerNumberOfItems[len(m.buttons)]; ok {
		for i, combinations := range allCombinations {
			for _, combination := range combinations {
				currentState := slices.Clone(m.initialState)
				for _, id := range combination {
					if actions, ok := m.buttons[id]; ok {
						for _, action := range actions {
							currentState[action] = toggleState(currentState[action])
						}
					}
				}
				if isOn(currentState, m.requiredState) {
					number := i + 2
					return &number, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("could not turn on machine with initial state %v and required state %v", m.initialState, m.requiredState)
}

func toMachine(data string) (*Machine, error) {
	inputs := strings.Split(data, " ")

	requiredStateData := inputs[0]
	initialStates, requiredStates := toStates(requiredStateData)

	buttonData := inputs[1 : len(inputs)-1]
	buttons, err := toButtons(buttonData)
	if err != nil {
		return nil, err
	}

	return &Machine{initialState: initialStates, requiredState: requiredStates, buttons: buttons}, nil
}

func toStates(data string) ([]bool, []bool) {
	input := strings.Split(data, "")
	states := []bool{}
	requiredStates := []bool{}
	for _, value := range input[1 : len(input)-1] {
		states = append(states, false)
		if value == "#" {
			requiredStates = append(requiredStates, true)
		} else {
			requiredStates = append(requiredStates, false)
		}
	}
	return states, requiredStates
}

func toButtons(data []string) (map[int][]int, error) {
	buttons := map[int][]int{}
	for i, input := range data {
		withoutPrefix := strings.TrimPrefix(input, "(")
		withoutSuffix := strings.TrimSuffix(withoutPrefix, ")")
		actions := []int{}
		values := strings.Split(withoutSuffix, ",")
		for _, value := range values {
			number, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			actions = append(actions, number)
		}
		buttons[i] = actions
	}
	return buttons, nil
}

func toggleState(state bool) bool {
	return !state
}

func isOn(state, requiredState []bool) bool {
	for i := 0; i < len(state); i++ {
		if state[i] != requiredState[i] {
			return false
		}
	}
	return true
}
