package main

import (
	"fmt"
)

type JunctionBox struct {
	id      int
	circuit int
}

type Storage struct {
	boxes map[int]*JunctionBox
}

func NewStorage(boxes map[int]*JunctionBox) *Storage {
	return &Storage{boxes: boxes}
}

func (s *Storage) Box(id int) (*JunctionBox, error) {
	if box, ok := s.boxes[id]; ok {
		return box, nil
	}
	return nil, fmt.Errorf("No box found with id %d", id)
}

func (b *JunctionBox) SetCircuit(circuit int) {
	b.circuit = circuit
}
