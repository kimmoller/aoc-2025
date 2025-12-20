package main

import (
	"strconv"
)

type Dial struct {
	position    int
	min         int
	max         int
	atZero      int
	crossedZero int
}

func NewDial(position int, min int, max int) Dial {
	return Dial{
		position,
		min,
		max,
		0,
		0,
	}
}

func (d *Dial) Turn(input string) error {
	direction, ticks, err := parseInput(input)
	if err != nil {
		return err
	}

	rotations := *ticks / 100
	d.crossedZero += rotations

	if *ticks%100 == 0 {
		return nil
	}

	crossedZero := false
	position := 0
	remainder := *ticks % 100
	if *direction == "L" {
		position, crossedZero = d.turnLeft(remainder)
	} else {
		position, crossedZero = d.turnRight(remainder)
	}

	if position == 0 {
		d.atZero++
		return nil
	}
	if crossedZero {
		d.crossedZero++
		return nil
	}

	return nil
}

func (d *Dial) turnLeft(remainder int) (int, bool) {
	crossedZero := false
	newPosition := d.position - remainder
	// Do not count crossing zero when we start at zero
	if newPosition < d.min && d.position != 0 {
		crossedZero = true
	}
	if newPosition < d.min {
		newPosition += 100
	}
	d.position = newPosition
	return newPosition, crossedZero
}

func (d *Dial) turnRight(remainder int) (int, bool) {
	crossedZero := false
	newPosition := d.position + remainder
	// Do not count crossing zero when we start at zero
	if newPosition > d.max && d.position != 0 {
		crossedZero = true
	}
	if newPosition > d.max {
		crossedZero = true
		newPosition -= 100
	}
	d.position = newPosition
	return newPosition, crossedZero
}

func parseInput(input string) (*string, *int, error) {
	direction := input[:1]
	ticksStr := input[1:]
	ticks, err := strconv.Atoi(ticksStr)
	if err != nil {
		return nil, nil, err
	}
	return &direction, &ticks, nil
}
