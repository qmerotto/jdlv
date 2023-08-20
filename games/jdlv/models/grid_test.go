package models

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCopy(t *testing.T) {
	oldGrid := NewGrid(10, 10)
	oldGrid[0][0].State = CellState{
		Alive:       true,
		Temperature: 35,
		Fuel:        10,
	}
	oldGrid[4][5].State = CellState{
		Alive:       false,
		Temperature: 35,
		Fuel:        10,
	}
	oldGrid[7][7].State = CellState{
		Alive:       true,
		Temperature: 0,
		Fuel:        10,
	}
	newGrid := oldGrid.copy()

	newGrid.Print()
	assert.Equal(t, oldGrid[0][0], newGrid[0][0])
	assert.Equal(t, oldGrid[4][5], newGrid[4][5])
	assert.Equal(t, oldGrid[7][7], newGrid[7][7])
}
