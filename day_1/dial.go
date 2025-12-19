package main

type Dial struct {
	position int
	min      int
	max      int
}

func NewDial(position int, min int, max int) Dial {
	return Dial{
		position,
		min,
		max,
	}
}

func (d *Dial) Position() int {
	return d.position
}

func (d *Dial) TurnLeft(ticks int) int {
	rotations := ticks / 100
	if ticks%100 == 0 {
		return rotations
	}
	for i := ticks; i > 100; i -= 100 {
		if ticks < 100 {
			newPosition := d.position - ticks
			if newPosition < d.min {
				newPosition += 100
			}
			d.position = newPosition
		}
	}
	return rotations
}

func (d *Dial) TurnRight(ticks int) int {
	rotations := ticks / 100
	if ticks%100 == 0 {
		return rotations
	}
	for i := ticks; i > 100; i -= 100 {
		if ticks < 100 {
			newPosition := d.position + ticks
			if newPosition > d.max {
				newPosition -= 100
			}
			d.position = newPosition
		}
	}
	return rotations
}
