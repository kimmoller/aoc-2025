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

func (p *PowerSupply) MaximumJoltage() (*int, error) {
	maxJoltage := 0
	for _, bank := range p.banks {
		value, err := joltage(bank)
		if err != nil {
			return nil, err
		}
		maxJoltage += *value
	}

	return &maxJoltage, nil
}

func joltage(bank BatteryBank) (*int, error) {
	firstValueIndex := 0
	firstValue := 0
	secondValue := 0

	batteries := bank.batteries
	for i := 0; i < len(bank.batteries)-1; i++ {
		battery := batteries[i]
		if battery.joltage > firstValue {
			firstValue = battery.joltage
			firstValueIndex = i
		}
	}

	for i := firstValueIndex + 1; i < len(bank.batteries); i++ {
		battery := batteries[i]
		if battery.joltage > secondValue {
			secondValue = battery.joltage
		}
	}

	strJoltage := strconv.Itoa(firstValue) + strconv.Itoa(secondValue)

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
