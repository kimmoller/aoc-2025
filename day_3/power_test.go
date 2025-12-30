package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoltageWithTwoBatteries(t *testing.T) {
	bank := BatteryBank{
		[]Battery{{9}, {8}, {7}, {6}, {5}, {4}, {3}, {2}, {1}, {1}, {1}, {1}, {1}, {1}, {1}},
	}
	bank2 := BatteryBank{
		[]Battery{{8}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {9}},
	}
	bank3 := BatteryBank{
		[]Battery{{2}, {3}, {4}, {2}, {3}, {4}, {2}, {3}, {4}, {2}, {3}, {4}, {2}, {7}, {8}},
	}
	bank4 := BatteryBank{
		[]Battery{{8}, {1}, {8}, {1}, {8}, {1}, {9}, {1}, {1}, {1}, {1}, {2}, {1}, {1}, {1}},
	}
	value, err := joltage(bank, 2)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 98, *value)

	value, err = joltage(bank2, 2)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 89, *value)

	value, err = joltage(bank3, 2)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 78, *value)

	value, err = joltage(bank4, 2)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 92, *value)
}

func TestJoltageWithTwelveBatteries(t *testing.T) {
	bank := BatteryBank{
		[]Battery{{9}, {8}, {7}, {6}, {5}, {4}, {3}, {2}, {1}, {1}, {1}, {1}, {1}, {1}, {1}},
	}
	bank2 := BatteryBank{
		[]Battery{{8}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}, {9}},
	}
	bank3 := BatteryBank{
		[]Battery{{2}, {3}, {4}, {2}, {3}, {4}, {2}, {3}, {4}, {2}, {3}, {4}, {2}, {7}, {8}},
	}
	bank4 := BatteryBank{
		[]Battery{{8}, {1}, {8}, {1}, {8}, {1}, {9}, {1}, {1}, {1}, {1}, {2}, {1}, {1}, {1}},
	}
	value, err := joltage(bank, 12)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 987654321111, *value)

	value, err = joltage(bank2, 12)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 811111111119, *value)

	value, err = joltage(bank3, 12)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 434234234278, *value)

	value, err = joltage(bank4, 12)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 888911112111, *value)
}
