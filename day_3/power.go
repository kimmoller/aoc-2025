package main

import (
	"strconv"
	"strings"
)

type PowerSupply struct {
	banks []BatteryBank
}

type BatteryBank struct {
	batteries []Battery
}

type Battery struct {
	joltage int
}

func NewPowerSupply(data []string) (*PowerSupply, error) {
	banks, err := batteryBanks(data)
	if err != nil {
		return nil, err
	}
	return &PowerSupply{banks}, nil
}

func (p *PowerSupply) MaximumJoltage(amountToActivate int) (*int, error) {
	maxJoltage := 0
	for _, bank := range p.banks {
		value, err := joltage(bank, amountToActivate)
		if err != nil {
			return nil, err
		}
		maxJoltage += *value
	}

	return &maxJoltage, nil
}

func joltage(bank BatteryBank, amountToActivate int) (*int, error) {
	joltages := []int{}

	batteries := bank.batteries
	numberOfBatteries := len(batteries)
	latestIndex := 0

	for i := 0; i < amountToActivate; i++ {
		highestValue := 0
		highestValueIndex := 0

		leftToActivate := amountToActivate - len(joltages)
		loopLength := numberOfBatteries - (latestIndex + leftToActivate)
		// This is required to make sure that the first loop takes into consideration that
		// one battery will be activated on this run
		if i == 0 {
			loopLength++
		}

		if loopLength <= 0 {
			break
		}

		startingIndex := 0
		if i != 0 {
			startingIndex = latestIndex + 1
		}
		endingIndex := startingIndex + loopLength

		for j := startingIndex; j < endingIndex; j++ {
			battery := batteries[j]
			if battery.joltage > highestValue {
				highestValue = battery.joltage
				highestValueIndex = j
			}
		}
		joltages = append(joltages, highestValue)
		latestIndex = highestValueIndex
	}

	if len(joltages) < amountToActivate {
		loopLength := amountToActivate - len(joltages)
		for i := 0; i < loopLength; i++ {
			indexToPick := i + latestIndex + 1
			joltages = append(joltages, batteries[indexToPick].joltage)
		}
	}

	strJoltage := ""
	for _, value := range joltages {
		strJoltage += strconv.Itoa(value)
	}

	joltage, err := strconv.Atoi(strJoltage)
	if err != nil {
		return nil, err
	}

	return &joltage, nil
}

func batteryBanks(data []string) ([]BatteryBank, error) {
	banks := []BatteryBank{}
	for _, input := range data {
		batteries := []Battery{}
		singleValues := strings.Split(input, "")
		for _, singleValue := range singleValues {
			value, err := strconv.Atoi(singleValue)
			if err != nil {
				return nil, err
			}
			battery := Battery{value}
			batteries = append(batteries, battery)
		}
		bank := BatteryBank{batteries}
		banks = append(banks, bank)
	}
	return banks, nil
}
