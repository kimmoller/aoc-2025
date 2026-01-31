package main

import (
	"aoc2025/utils"
)

func Run(path string) (*int, error) {
	data, err := utils.ReadData(path)
	if err != nil {
		return nil, err
	}

	machines, mostButtons, err := machinesFrom(data)
	if err != nil {
		return nil, err
	}

	combinations, err := AllCombinations(*mostButtons)
	if err != nil {
		return nil, err
	}

	sum := 0
	for _, machine := range machines {
		requiredPresses, err := machine.TurnOn(combinations)
		if err != nil {
			return nil, err
		}
		sum += *requiredPresses
	}

	return &sum, nil
}

func machinesFrom(data []string) ([]*Machine, *int, error) {
	machines := []*Machine{}
	mostButtons := 0
	for _, input := range data {
		machine, err := NewMachine(input)
		if err != nil {
			return nil, nil, err
		}
		if mostButtons < len(machine.buttons) {
			mostButtons = len(machine.buttons)
		}
		machines = append(machines, machine)
	}
	return machines, &mostButtons, nil
}
