package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBeam(t *testing.T) {
	teleporter := NewTeleporter([]string{
		".......S.......",
		"...............",
		".......^.......",
	})
	beam := NewBeam(0, 7)
	newBeams := beam.Run(teleporter)
	assert.Equal(t, 2, len(newBeams))
}
