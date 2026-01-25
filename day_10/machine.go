package main

import (
	"strconv"
	"strings"
)

type Machine struct {
	currentState  []bool
	requiredState []bool
	buttons       []Button
}

type Button struct {
	actions []int
}

func NewMachine(data string) (*Machine, error) {
	return toMachine(data)
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

	return &Machine{currentState: initialStates, requiredState: requiredStates, buttons: buttons}, nil
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

func toButtons(data []string) ([]Button, error) {
	buttons := []Button{}
	for _, input := range data {
		withoutPrefix := strings.TrimPrefix(input, "(")
		withoutSuffix := strings.TrimSuffix(withoutPrefix, ")")
		button := Button{}
		values := strings.Split(withoutSuffix, ",")
		for _, value := range values {
			number, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			button.actions = append(button.actions, number)
		}
		buttons = append(buttons, button)
	}
	return buttons, nil
}
