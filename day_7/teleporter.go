package main

import (
	"strconv"
	"strings"
)

const (
	START       = "S"
	EMPTY_SPACE = "."
	SPLITTER    = "^"
)

type Teleporter struct {
	manifold       map[int][]string
	beams          []Beam
	splitPositions map[string]struct{}
}

func NewTeleporter(data []string) *Teleporter {
	teloporter := Teleporter{beams: []Beam{}, splitPositions: map[string]struct{}{}}
	teloporter.FillManifold(data)
	return &teloporter
}

func (t *Teleporter) Height() int {
	return len(t.manifold)
}

func (t *Teleporter) Beams() []Beam {
	return t.beams
}

func (t *Teleporter) FillManifold(data []string) {
	manifold := map[int][]string{}
	for i, row := range data {
		items := strings.Split(row, "")
		manifold[i] = items
	}
	t.manifold = manifold
}

func (t *Teleporter) Start() {
	manifold := t.manifold
	firstRow := manifold[0]
	startingIndex := 0
	for i, item := range firstRow {
		if item == START {
			startingIndex = i
		}
	}

	firstBeam := NewBeam(0, startingIndex)
	t.RunBeams([]Beam{*firstBeam})
}

func (t *Teleporter) RunBeams(beams []Beam) {
	t.beams = append(t.beams, beams...)
	for _, beam := range beams {
		newBeams := beam.Run(t)
		if len(newBeams) != 0 {
			t.RunBeams(newBeams)
		}
	}
}

func (t *Teleporter) AddSplitPosition(row, position int) {
	rowStr := strconv.Itoa(row)
	positionStr := strconv.Itoa(position)
	index := rowStr + positionStr
	t.splitPositions[index] = struct{}{}
}
