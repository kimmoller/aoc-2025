package main

type Beam struct {
	startingRow   int
	startingIndex int
}

func NewBeam(row int, index int) *Beam {
	return &Beam{startingRow: row, startingIndex: index}
}

func (b *Beam) Run(teleporter *Teleporter) []Beam {
	startingRow := b.startingRow + 1
	position := b.startingIndex
	height := teleporter.Height()
	maxDrops := height - startingRow

	for i := 0; i < maxDrops; i++ {
		actualRow := i + startingRow
		if row, ok := teleporter.manifold[actualRow]; ok {
			nextPosition := row[position]
			if nextPosition == SPLITTER {
				teleporter.AddSplitPosition(actualRow, position)
				newLeftBeam := NewBeam(actualRow, position-1)

				// Only create a new right beam if there won't be a splitter making a left beam on the same positon
				twoOver := position + 2
				if len(row) <= twoOver || row[twoOver] != SPLITTER {
					newRightBeam := NewBeam(actualRow, position+1)
					return []Beam{*newLeftBeam, *newRightBeam}
				}
				return []Beam{*newLeftBeam}
			}
		}
	}
	return []Beam{}
}
